# Day 8

This day was particularly challenging for me but I was greatly helped by an example I found on reddit here: https://github.com/kevin-schmid/AdventOfCode/blob/main/2024/08/p2/main.go

My main issue was how I was handling the out of bounds checks as I was setting the first iteration to always be true and this was giving me the answer 1 too high, using this persons example I changed my check to do one before going in.

I also liked the idea of having a .toString() method on the position type and using this to track and prevent counting overlaps twice so I changed my original implementation of copyingthe input matrix to track antinodes by placing #'s.
