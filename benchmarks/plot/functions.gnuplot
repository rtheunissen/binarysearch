
log2(x) = log(x)/log(2)

mkdir(file) = sprintf("mkdir -p $(dirname %s)", file)

alpha(i) = (int(255.*i) << 24) + 0xEF8000

title(t) = sprintf("set label '%s' at screen 0.5,screen 0.98 font ',20' center back", t)