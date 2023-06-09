log2(x) = log(x)/log(2)

mkdir(file) = sprintf("mkdir -p $(dirname %s)", file)

filter(col, iff, val) = stringcolumn(col) eq (iff) ? (val) : (NaN)

