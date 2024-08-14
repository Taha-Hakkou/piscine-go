#!/usr/bin/bash

# Find clues:
#$ grep "CLUE" crimescene
#> CLUE: Footage from an ATM security camera is blurry but shows that the perpetrator is a tall male, at least 6'.
#> CLUE: Found a wallet believed to belong to the killer: no ID, just loose change, and membership cards for AAA, Delta SkyMiles, the local library, and the Museum of Bash History. The cards are totally untraceable and have no name, for some reason.
#> CLUE: Questioned the barista at the local coffee shop. He said a woman left right before they heard the shots. The name on her latte was Annabel, she had blond spiky hair and a New Zealand accent.

# Based on clue3, find Annabel's address:
#$ grep "Annabel" people | grep "	F"
#> Annabel Sun	F	26	Hart Place, line 40
#> Annabel Church	F	38	Buckingham Place, line 179

# Check lines on streets:
#$ head -n 40 mystery/streets/Hart_Place | tail -n 1
#> SEE INTERVIEW #47246024
#$ head -n 179 mystery/streets/Buckingham_Place | tail -n 1
#> SEE INTERVIEW #699607

# Check interviews:
#$ cat interviews/interview-47246024
#> Ms. Sun has brown hair and is not from New Zealand.  Not the witness from the cafe.
#$ cat interviews/interview-699607
#> Interviewed Ms. Church at 2:04 pm.  Witness stated that she did not see anyone she could identify as the shooter, that she ran away as soon as the shots were fired.
#However, she reports seeing the car that fled the scene.  Describes it as a blue Honda, with a license plate that starts with "L337" and ends with "9"

# Find owners of cars matching Annabel's description and clue1:
#$ grep -A 5 L337 mystery/vehicles | grep -A 4 Honda | grep -A 3 Blue | grep -B 1 6\' | grep Owner
#> Owner: Erika Owens
#> Owner: Joe Germuska
#> Owner: Dartey Henv
#> Owner: Hellen Maher

# Check if has all found memberships (from clue2): (must return 4)
#$ cat AAA Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History | grep -c "Erika Owens"
#> 0
#$ cat AAA Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History | grep -c "Joe Germuska"
#> 2
#$ cat AAA Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History | grep -c "Dartey Henv"
#> 4
#$ cat AAA Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History | grep -c "Hellen Maher"
#> 4

# Hence the perpetrator is a Male (clue1):
echo "Dartey Henv"
