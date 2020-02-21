from mcr.microsoft.com/windows/nanoserver:1809
# from mcr.microsoft.com/windows/servercore:1809

copy ./domain-socket-proxy.exe /domain-socket-proxy.exe

ENTRYPOINT [ "/domain-socket-proxy.exe", "serve"]

USER ContainerAdministrator