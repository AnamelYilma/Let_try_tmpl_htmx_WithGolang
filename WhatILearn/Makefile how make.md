
# How To Install and ready Make
## For Frist Time 
### install on admin powershell by run this
        
        choco install make
            
and this  

        Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
### For ever project 
create make file 
    run by 
        make run

to create a Makefile
create Makefile
then write how to run you code 
           

        TEMPL := templ
        GO := go

        .PHONY: generate run build tidy

        generate:
            $(TEMPL) generate

        run: generate
            $(GO) run main.go

        build: generate
            $(GO) build .

        tidy:
            $(GO) mod tidy

    

