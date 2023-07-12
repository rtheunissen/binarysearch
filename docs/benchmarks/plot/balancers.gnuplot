load "docs/benchmarks/plot/colors.gnuplot"

##################################################################
#
#           BALANCERS
#
##################################################################

Median  = 1
Height  = 2
Weight  = 3
Log     = 4
Cost    = 5
DSW     = 6

set style line Median  dashtype 1 ps 1 lw 0.5 pt  11 pn 2 lc rgb COLOR_BLACK
set style line Height  dashtype 4 ps 1 lw 1   pt  10 pn 2 lc rgb COLOR_BROWN
set style line Weight  dashtype 1 ps 1 lw 1   pt   9 pn 2 lc rgb COLOR_BLUE
set style line Log     dashtype 5 ps 1 lw 1   pt   8 pn 2 lc rgb COLOR_CYAN
set style line Cost    dashtype 2 ps 1 lw 1   pt  14 pn 2 lc rgb COLOR_GREEN
set style line DSW     dashtype 3 ps 1 lw 1   pt   6 pn 2 lc rgb COLOR_PINK

# set style line HalfSize    dashtype 1 ps 1 lw 1 pt   9 pn 2 lc rgb COLOR_RED
# set style line LogSize     dashtype 5 ps 1 lw 1 pt   8 pn 2 lc rgb COLOR_YELLOW
# set style line HalfWeight  dashtype 1 ps 1 lw 1 pt  13 pn 2 lc rgb COLOR_BLUE
# set style line LogWeight   dashtype 5 ps 1 lw 1 pt  12 pn 2 lc rgb COLOR_CYAN
# set style line Cost        dashtype 1 ps 1 lw 1 pt  14 pn 2 lc rgb COLOR_BROWN
# set style line DSW         dashtype 3 ps 1 lw 1 pt   6 pn 2 lc rgb COLOR_PINK

##################################################################
#
#           GROUPS
#
##################################################################

Partition = "Median Height Weight Log Cost"
All       = "Median Height Weight Log Cost DSW"

GROUPS = "Partition All"

DISTRIBUTIONS = "Uniform"

do for [GROUP in GROUPS] {

    ##################################################################
    #
    #           PARTITION COUNT
    #
    ##################################################################

    MEASUREMENT = "PartitionCount"

    set xlabel "{Size × 10^5}"
    set ylabel "{/:Bold Partition Count} / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionCount')/column('Size')"

    set format y2 "%.2f"

    DATA = "docs/benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           TOTAL PARTITION DEPTH
    #
    ##################################################################

    MEASUREMENT = "TotalPartitionDepth"

    set xlabel "{Size × 10^5}"
    set ylabel "{/:Bold Partition Depth} / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionDepth')/column('Size')"

    set format y2 "%.2f"

    DATA = "docs/benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"


    ##################################################################
    #
    #           AVERAGE PARTITION DEPTH
    #
    ##################################################################

    MEASUREMENT = "AveragePartitionDepth"

    set xlabel "{Size × 10^5}"
    set ylabel "{/:Bold Partition Depth } / Partition Count"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionDepth')/column('PartitionCount')"

    set format y2 "%.2f"

    DATA = "docs/benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           ROTATIONS
    #
    ##################################################################

    MEASUREMENT = "Rotations"

    set xlabel "{Size × 10^5}"
    set ylabel "{/:Bold Rotations}"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Rotations')"

    set format y2 "%.0f"

    DATA = "docs/benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"


    ##################################################################
    #
    #           MAXIMUM PATH LENGTH
    #
    ##################################################################

    MEASUREMENT = "MaximumPathLength"

    set xlabel "Size × 10^5"
    set ylabel "{/:Bold Maximum Path Length } / log_2(n)"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('MaximumPathLength')/log2(column('Size'))"

    set format y2 "%.2f"

    DATA = "docs/benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"


    ##################################################################
    #
    #           AVERAGE PATH LENGTH
    #
    ##################################################################

    MEASUREMENT = "AveragePathLength"

    set xlabel "Size × 10^5"
    set ylabel "{/:Bold Average Path Length } / log_2(n)"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('AveragePathLength')/log2(column('Size'))"

    set format y2 "%.2f"

    DATA = "docs/benchmarks/data/balancers/measurements"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           DURATION
    #
    ##################################################################

    MEASUREMENT = "Duration"

    set xlabel "Size × 10^6"
    set ylabel "{/:Bold Duration } in milliseconds / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Duration')/column('Size')"

    set format y2 "%.0f"

    DATA = "docs/benchmarks/data/balancers/benchmarks"

    SMOOTH = "sbezier"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    SMOOTH = "unique"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"

    ##################################################################
    #
    #           DURATION (CUMULATIVE)
    #
    ##################################################################

    MEASUREMENT = "TotalDuration"

    set xlabel "Size × 10^6"
    set ylabel "{/:Bold Total Duration } in seconds / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Duration')/column('Size')/1000"

    set format y2 "%.2f"

    DATA = "docs/benchmarks/data/balancers/benchmarks"
    SMOOTH = "cumulative"
    load "docs/benchmarks/plot/balancers/lines.gnuplot"
}
