##############################################################################
# 2018 Advent of Code
# Day: 1
# Part: 1
# Answer: 553
#
# @author Matt Adorjan
##############################################################################

# Read frequency changes in to a variable
with open('./input.txt', 'r') as f:
    data_in = f.readlines()

# Set initial frequecy as 0
FREQ = 0

# For each of the frequency changes, add to the current frequency value
for delta in data_in:
    FREQ = FREQ + int(delta.strip())    # Convert from str -> int and strip off line endings (\n)

print("Final frequency after calibration is: {}".format(FREQ))
