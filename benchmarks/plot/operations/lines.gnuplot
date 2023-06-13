load "benchmarks/plot/graph.gnuplot"
load "benchmarks/plot/functions.gnuplot"

set tmargin 9
set bmargin 5
set lmargin 5

SVG = sprintf("benchmarks/svg/operations/%s/%s/%s__%s.svg", OPERATION, GROUP, MEASUREMENT, SMOOTH)
system mkdir(SVG)
set output SVG
print SVG

set table $TABLE
plot for [STRATEGY in @GROUP] DATA.'/'.STRATEGY using (@x):(@y) smooth @SMOOTH
unset table

set title "{/:Bold ".OPERATION."}"

plot for [i=1:words(@GROUP)] $TABLE \
    index (i-1) \
    using 1:2 \
    axes x1y2 \
    smooth @SMOOTH \
    with linespoints \
    linestyle value(word(@GROUP,i)) \
    title word(@GROUP,i)

do for [DISTRIBUTION in DISTRIBUTIONS] {

    SVG = sprintf("benchmarks/svg/operations/%s/%s/%s/%s__%s.svg", OPERATION, GROUP, DISTRIBUTION, MEASUREMENT, SMOOTH)
    system mkdir(SVG)
    set output SVG
    print SVG

    set table $TABLE
    plot for [STRATEGY in @GROUP] DATA.'/'.STRATEGY using (@x):(filter('Distribution', DISTRIBUTION, @y)) smooth @SMOOTH
    unset table

    set title OPERATION." - ".DISTRIBUTION

    plot for [i=1:words(@GROUP)] $TABLE \
        index (i-1) \
        using 1:2 \
        axes x1y2 \
        smooth mcsplines \
        with linespoints \
        linestyle value(word(@GROUP,i)) \
        title word(@GROUP,i)
}