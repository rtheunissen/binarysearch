load "benchmarks/plot/graph.gnuplot"
load "benchmarks/plot/functions.gnuplot"

set tmargin 10
set bmargin 5
set lmargin 1

SVG = sprintf("benchmarks/svg/balancers/%s/%s__%s.svg", GROUP, MEASUREMENT, SMOOTH)
system mkdir(SVG)
set output SVG
print SVG

set table $TABLE
plot for [BALANCER in @GROUP] DATA.'/'.BALANCER using (@x):(@y) smooth @SMOOTH
unset table

plot for [i=1:words(@GROUP)] $TABLE \
    index (i-1) \
    using 1:2 \
    axes x1y2 \
    smooth @SMOOTH \
    with linespoints \
    linestyle value(word(@GROUP,i)) \
    title word(@GROUP,i)

do for [DISTRIBUTION in DISTRIBUTIONS] {

    SVG = sprintf("benchmarks/svg/balancers/%s/%s/%s__%s.svg", GROUP, DISTRIBUTION, MEASUREMENT, SMOOTH)
    system mkdir(SVG)
    set output SVG
    print SVG

    set table $TABLE
    plot for [BALANCER in @GROUP] DATA.'/'.BALANCER using (@x):(filter('Distribution', DISTRIBUTION, @y)) smooth sbezier
    unset table

    plot for [i=1:words(@GROUP)] $TABLE \
        index (i-1) \
        using 1:2 \
        axes x1y2 \
        smooth mcsplines \
        with linespoints \
        linestyle value(word(@GROUP,i)) \
        title word(@GROUP,i)
}
