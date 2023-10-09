package portscan

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/Ullaakut/nmap/v3"
    "github.com/fatih/color"
)
var red = color.New(color.FgHiRed, color.Bold)
func Port() {
    var url string
	
    red.Print("Digite o host: ")
    fmt.Scanln(&url)
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()

    scanner, err := nmap.NewScanner(
        ctx,
        nmap.WithTargets(url),
        nmap.WithPorts("80,443,22,3306,1433,5432,6379,25,110,143,993,995,53,161,162,5900"),
    )
    if err != nil {
        log.Fatalf("unable to create nmap scanner: %v", err)
    }

    result, warnings, err := scanner.Run()
    if len(*warnings) > 0 {
        log.Printf("run finished with warnings: %s\n", *warnings)
    }
    if err != nil {
        log.Fatalf("unable to run nmap scan: %v", err)
    }

    green := color.New(color.FgGreen).SprintFunc()
    red := color.New(color.FgHiRed,color.Bold).SprintFunc()

    for _, host := range result.Hosts {
        if len(host.Ports) == 0 || len(host.Addresses) == 0 {
            continue
        }

        fmt.Printf("Host %q:\n", host.Addresses[0])

        for _, port := range host.Ports {
            var portStatus string
            if portStatus == "open" {
                portStatus = green(port.State)
            } else {
                portStatus = red(port.State)
            }
            fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, portStatus, port.Service.Name)
        }
    }

    fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}
