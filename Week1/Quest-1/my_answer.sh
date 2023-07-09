#! /bin/bash 

# ---

# AAA, Delta SkyMiles, the local library, and the Museum of Bash History

# CLUE: Footage from an ATM security camera is blurry but shows that the perpetrator is a tall male, at least 6'.

# CLUE: Found a wallet believed to belong to the killer: no ID, just loose change, and membership cards for AAA, Delta SkyMiles, the local library, and the Museum of Bash History. The cards are totally untraceable and have no name, for some reason.

# CLUE: Questioned the barista at the local coffee shop. He said a woman left right before they heard the shots. The name on her latte was Annabel, she had blond spiky hair and a New Zealand accent.

# ---

# ---

# Vehicle and owner information from the Terminal City Department of Motor Vehicles

# ---

# (base) can@Cans-MacBook-Pro mystery % head -n 20 people | grep "Annabel" people
# Annabel Sun F 26 Hart Place, line 40
# Oluwasegun Annabel M 37 Mattapan Street, line 173
# Annabel Church F 38 Buckingham Place, line 179
# Annabel Fuglsang M 40 Haley Street, line 176

# (base) can@Cans-MacBook-Pro mystery % head -n 40 streets/Hart_Place | tail -n 1
# SEE INTERVIEW #47246024

# (base) can@Cans-MacBook-Pro mystery % head -n 179 streets/Buckingham_Place | tail -n 1
# SEE INTERVIEW #699607

# (base) can@Cans-MacBook-Pro mystery % head interviews/interview-699607 | tail -n 1
# However, she reports seeing the car that fled the scene. Describes it as a blue Honda, with a license plate that starts with "L337" and ends with "9"
# (base) can@Cans-MacBook-Pro mystery % head interviews/interview-47246024 | tail -n 1
# Ms. Sun has brown hair and is not from New Zealand. Not the witness from the cafe.

# grep -A 5 -B 8 'Color: Blue' vehicles | grep -A 5 -B 2 'Make: Honda' | grep -A 5 -B 3 '^License Plate L337.\*9$' > filtered_vehicles.txt

# grep -B 5 "Height: 6'" filtered_vehicles.txt

# (base) can@Cans-MacBook-Pro memberships % cat Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History AAA | grep -c "Dartey Henv"
# 4
# (base) can@Cans-MacBook-Pro memberships % cat Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History AAA | grep -c "Hellen Maher"
# 4
# (base) can@Cans-MacBook-Pro memberships % cat Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History AAA | grep -c "Joe Germuska"
# 2
# (base) can@Cans-MacBook-Pro memberships % cat Delta_SkyMiles Terminal_City_Library Museum_of_Bash_History AAA | grep -c "Erika Owens"
# 0 

echo "Dartey Henv"
