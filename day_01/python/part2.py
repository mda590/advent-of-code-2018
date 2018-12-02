##############################################################################
# 2018 Advent of Code
# Day: 1
# Part: 2
# Answer: 78724
#
# @author Matt Adorjan
##############################################################################

# Read frequency changes in to a variable
with open('../input.txt', 'r') as f:
    data_in = f.readlines()

# Set initial frequecy as 0
FREQ = 0
FREQ_LIST = [FREQ]
CHANGE_IND = 0

# For each of the frequency changes, add to the current frequency value
while 1:
    FREQ = FREQ + int(data_in[CHANGE_IND].strip())    # Convert from str -> int and strip off line endings (\n)

    if FREQ in FREQ_LIST:
        print("The first frequency to appear twice is: {}".format(FREQ))
        break
    
    FREQ_LIST.append(FREQ)

    CHANGE_IND = CHANGE_IND + 1

    if CHANGE_IND == len(data_in):
        CHANGE_IND = 0
