#!/usr/bin/python

from itertools import permutations

data = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b"]

counter = 0
for p in permutations(data):
    #print p
    counter = counter + 1

print counter