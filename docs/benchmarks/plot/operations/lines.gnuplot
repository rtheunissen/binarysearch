load "docs/benchmarks/plot/graph.gnuplot"
load "docs/benchmarks/plot/functions.gnuplot"

set tmargin 5
set bmargin 3
set lmargin 4
set rmargin 4

set title offset 0,2.5
set key at graph 0.5, screen 0.85

SVG = sprintf("docs/benchmarks/svg/operations/%s/%s/%s__%s.svg", GROUP, OPERATION, MEASUREMENT, SMOOTH)
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
    with linespoints \
    linestyle value(word(@GROUP,i)) \
    title word(@GROUP,i)

do for [DISTRIBUTION in DISTRIBUTIONS] {

    SVG = sprintf("docs/benchmarks/svg/operations/%s/%s/%s/%s__%s.svg", GROUP, OPERATION, DISTRIBUTION, MEASUREMENT, SMOOTH)
    system mkdir(SVG)
    set output SVG
    print SVG

    set table $TABLE
    plot for [STRATEGY in @GROUP] DATA.'/'.STRATEGY using (@x):(filter('Distribution', DISTRIBUTION, @y)) smooth @SMOOTH
    unset table

    set title "{/:Bold ".OPERATION."} — {/:Italic ".DISTRIBUTION."}"

    plot for [i=1:words(@GROUP)] $TABLE \
        index (i-1) \
        using 1:2 \
        axes x1y2 \
        with linespoints \
        linestyle value(word(@GROUP,i)) \
        title word(@GROUP,i)
}