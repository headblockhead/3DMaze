@echo off
mkdir "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3DMAZE\"
mkdir "C:\3DMAZE\"
copy arcadesim.exe C:\3DMAZE\
copy textures.png  C:\3DMAZE\
move "arcade sim.lnk" "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3DMAZE"
