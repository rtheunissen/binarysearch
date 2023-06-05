load "benchmarks/plot/style.gnuplot"
load "benchmarks/plot/functions.gnuplot"

set tmargin 8

set bmargin 4

svg = "benchmarks/svg/balancers/".GROUP."/".MEASUREMENT."__".smooth."__lines.svg"

system mkdir(svg)

set output svg

print svg

set table $TEMP
eval "plot for [i=1:|".GROUP."|] csv_dir.'/'.".GROUP."[i] using (".x."):(".y.") smooth ".smooth
unset table
eval "plot for [i=1:|".GROUP."|] $TEMP index (i-1) using 1:2 smooth ".smooth." axes x1y2 with lp ls value(".GROUP."[i]) title ".GROUP."[i]"

do for [Distribution=1:|Distributions|] {

    DISTRIBUTION = Distributions[Distribution]

    svg = "benchmarks/svg/balancers/".GROUP."/".DISTRIBUTION."/".MEASUREMENT."__".smooth."__lines.svg"

    system mkdir(svg)

    set output svg

    print svg

    set table $TEMP

    eval "plot for [i=1:|".GROUP."|] csv_dir.'/'.".GROUP."[i] using (".x."):(stringcolumn('Distribution') eq '".DISTRIBUTION."' ? (".y.") : NaN) smooth ".smooth

    unset table

    eval "plot for [i=1:|".GROUP."|] $TEMP index (i-1) using 1:2 smooth ".smooth." axes x1y2 with lp linestyle value(".GROUP."[i]) title ".GROUP."[i]"
}
