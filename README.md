# Go8 a simple Chip-8 emulator written in golang   
----------

This emulator is a little project of mine I made to learn more about low level concepts of computers (yes I know this machine is technically a VM but still)   
----------

## Dependencies   
   
The only library this application depends on is the RayLibGo bindings so be sure you have raylib installed  

----------

## How to use?   
---------
To load a rom simply provide it into the programs command line arguments   
The key board is mapped to the left side of the keyboard   
1,2,3,C    
(1),(2),(3),(4)        
4,5,6,D    
(Q),(W),(E),(R)     
7,8,9,E    
(A),(S),(D),(F)       
A,0,B,F    
(Z),(X),(C),(V)          
To change the target frame rate(instructions\s) press CTRL-F (it's kinda rudimentary right now but it works okay)     
## TODO
   
Optimize in places. I feel like this program could 100% be faster   
   
Implement Sound. It really won't be that hard to do but I'm getting bored of this project   
   
Maybe add a GUI interface for loading ROMS on the fly. Would be nice so you don't have to open a terminal to open ROMS   
