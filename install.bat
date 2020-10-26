echo Installing, This might take a few seconds...
mkdir "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\"
mkdir "C:\3DMAZE\"
copy ARCADE~1.exe C:\3DMAZE\
copy textures.png  C:\3DMAZE\
copy UNINST~1.exe  C:\3DMAZE\
copy Manual.txt  C:\3DMAZE\
move "Launch.lnk" "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\"
move "UNINST~1.lnk" "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\"
move "Manual.lnk" "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\3D Maze\"
pause