package main

import (
    "context"
    "crypto/md5"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

    "golang.org/x/sync/errgroup"
    "github.com/spf13/pflag"
)


func main() {
    var r *int = pflag.Int("routines", 3, "routines to run")
    pflag.Parse()

    m, err := MD5All(context.Background(), *r, ".")
    if err != nil {
        log.Fatal(err)
    }

    for k, sum := range m {
        fmt.Printf("%s:\t%x\n", k, sum)
    }
}

type result struct {
    path string
    sum [md5.Size]byte
}

func MD5All(ctx context.Context, routines int, root string) (map[string][md5.Size]byte, error) {
    g, ctx := errgroup.WithContext(ctx)
    paths := make(chan string)
    
    g.Go(func() error {
        defer close(paths)
        return filepath.Walk(root, func(path string, info os.FileInfo, err error) error  {
            if err != nil {
                return err
            }
            if !info.Mode().IsRegular() {
                return nil
            }
            select {
                case paths <- path:
                case <- ctx.Done():
                    return ctx.Err()
            }
            return nil
        })
    }) 

    // number of goroutines to read files
    c := make(chan result)
    for i := 0; i < routines; i++ {
        fmt.Println("starting reading worker thread...")
        g.Go(func() error {
            for path := range paths {
                data, err := ioutil.ReadFile(path)
                if err != nil {
                    return err
                }
                select {
                    case c <- result{path, md5.Sum(data)}:
                    case <- ctx.Done():
                        return ctx.Err()
                }
            }
            return nil
        })
    }

    go func() {
        g.Wait() // catch goroutine termination to close channel
        close(c)
    }()

    m := make(map[string][md5.Size]byte)
    for r := range c {
        m[r.path] = r.sum
    }

    err := g.Wait()
    if err != nil {
        return nil, err
    }

    return m, nil
}
