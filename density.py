import glob, os
os.chdir("instances")
for file in glob.glob("*.col"):
    with open(file, 'r') as rf:
      for line in rf.readlines():
        line = line.replace('\n','')
        if (line[0] == 'p'):
          tokens = line.split(' ')
          n, e = int(tokens[2]), int(tokens[3])
          print('{0:s}\t\t{1:.3f}%'.format(file, e*100/(n*n)))