load "benchmarks/plot/graph.gnuplot"
load "benchmarks/plot/functions.gnuplot"

set tmargin 3
set bmargin 3
set rmargin 4
set lmargin 4

SVG = sprintf("benchmarks/svg/balancers/%s/%s__%s.svg", GROUP, MEASUREMENT, SMOOTH)
system mkdir(SVG)
set output SVG
print SVG

plot for [BALANCER in @GROUP] DATA.'/'.BALANCER \
    using (@x):(@y) \
    axes x1y2 \
    smooth @SMOOTH \
    with linespoints \
    linestyle value(BALANCER) \
    title BALANCER

do for [DISTRIBUTION in DISTRIBUTIONS] {

    SVG = sprintf("benchmarks/svg/balancers/%s/%s/%s__%s.svg", GROUP, DISTRIBUTION, MEASUREMENT, SMOOTH)
    system mkdir(SVG)
    set output SVG
    print SVG

    plot for [BALANCER in @GROUP] DATA.'/'.BALANCER \
        using (@x):(@y) \
        axes x1y2 \
        smooth @SMOOTH \
        with linespoints \
        linestyle value(BALANCER) \
        title BALANCER
}
