# LEM-IN
## Karl Heinrich Tamm / kartamm

This project is a digital version of an ant farm. The ant farm composes of tunnels and rooms. Ants need to find their way from the starting room to the end-room turn-by-turn (moving rooms every turn), with the caveat that each room can only accommodate one ant at a time.

A Thank-you for support and inspiration to Tomi Kivilo!

# Auditing

Exercise description here: [description](https://01.kood.tech/git/root/public/src/branch/master/subjects/lem-in)

Audit checklist here: [audit](https://01.kood.tech/git/root/public/src/branch/master/subjects/lem-in/audit)

To make sure that an up-to-date version of the project is in use, run a "git pull" on the commandline

Run the test script test.sh to run tests semi-automatically with examples from the audit resources
```
bash test.sh  
or  
./test.sh  
if the above commands to not work, run  
chmod +x test.sh  
beforehand
```
or run commands manually  
```
go run . <path-to-file>
```
as in:  
```
  go run . example01.txt
```

### Thank you for taking the time to audit!!! Farewell! ###