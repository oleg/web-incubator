import png

s = ['110010010011',
     '101011010100',
     '110010110101',
     '100010010011']
s = [[int(c) for c in row] for row in s]

w = png.Writer(len(s[0]), len(s), greyscale=True, bitdepth=1)
f = open('png.png', 'wb')
w.write(f, s)
f.close()

p = [(255, 0, 0, 0, 255, 0, 0, 0, 255),
     (128, 0, 0, 0, 128, 0, 0, 0, 128)]
f = open('swatch.png', 'wb')
w = png.Writer(3, 2, greyscale=False)
w.write(f, p)
f.close()
