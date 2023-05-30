load "benchmarks/plot/style.gnuplot"
load "benchmarks/plot/functions.gnuplot"

# set title "{/:Bold ".OPERATION."}"

data = data_dir."/".OPERATION

svg = "benchmarks/svg/balancers/".OPERATION."/".GROUP."/".MEASUREMENT."__".smooth."__dots.svg"

#system mkdir(svg)

set output svg

print svg

set key samplen 1
set key width 3
set key at screen 0.55, screen 0.85


eval "plot for [i=1:|".GROUP."|] '".data."' index ".GROUP."[i] using (".x."):(".y.") smooth ".smooth." axes x1y2 with dots ls value(".GROUP."[i]) notitle, \
           for [i=1:|".GROUP."|] '".data."' index ".GROUP."[i] using (".x."):(".y.") smooth ".smooth." axes x1y2 with lp   ls value(".GROUP."[i]) lw 0 title ".GROUP."[i]"

do for [Distribution=1:|Distributions|] {

    DISTRIBUTION = Distributions[Distribution]

    # set title "{/:Bold ".OPERATION."} — {/:Italic ".DISTRIBUTION."}"

    svg = "benchmarks/svg/balancers/".OPERATION."/".GROUP."/".DISTRIBUTION."/".MEASUREMENT."__".smooth."__dots.svg"

    #system mkdir(svg)

    set output svg

    print svg

    eval "plot for [i=1:|".GROUP."|] '".data."' index ".GROUP."[i] using (".x."):(stringcolumn('Distribution') eq '".DISTRIBUTION."' ? (".y.") : NaN) smooth ".smooth." axes x1y2 with dots ls value(".GROUP."[i]) notitle, \
               for [i=1:|".GROUP."|] '".data."' index ".GROUP."[i] using (".x."):(stringcolumn('Distribution') eq '".DISTRIBUTION."' ? (".y.") : NaN) smooth ".smooth." axes x1y2 with   lp ls value(".GROUP."[i]) lw 0 title ".GROUP."[i]"
}

