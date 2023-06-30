load "benchmarks/plot/colors.gnuplot"

##################################################################
#
#           STRATEGIES
#
##################################################################

AVLBottomUp                = 1
AVLTopDown                 = 2
AVLWeakTopDown             = 3
AVLWeakBottomUp            = 4
AVLRelaxedTopDown          = 5
AVLRelaxedBottomUp         = 6
RedBlackBottomUp           = 7
RedBlackTopDown            = 8
RedBlackRelaxedBottomUp    = 9
RedBlackRelaxedTopDown     = 10
LBSTBottomUp               = 11
LBSTTopDown                = 12
LBSTRelaxed                = 13
TreapTopDown               = 14
TreapFingerTree            = 15
Randomized                 = 16
Zip                        = 17
Splay                      = 18
Conc                       = 19





set style line AVLBottomUp                  dashtype 1 ps 1 lw 1 pt   1 pn 2 lc rgb COLOR_BLACK
set style line AVLTopDown                   dashtype 3 ps 1 lw 1 pt   2 pn 2 lc rgb COLOR_CYAN
set style line AVLWeakBottomUp              dashtype 1 ps 1 lw 1 pt   4 pn 2 lc rgb COLOR_BLUE
set style line AVLWeakTopDown               dashtype 5 ps 1 lw 1 pt   3 pn 2 lc rgb COLOR_GREEN
set style line AVLRelaxedTopDown            dashtype 5 ps 1 lw 1 pt   6 pn 2 lc rgb COLOR_BLUE
set style line AVLRelaxedBottomUp           dashtype 1 ps 1 lw 1 pt   7 pn 2 lc rgb COLOR_YELLOW
set style line RedBlackBottomUp             dashtype 1 ps 1 lw 1 pt  12 pn 2 lc rgb COLOR_RED
set style line RedBlackTopDown              dashtype 3 ps 1 lw 1 pt   5 pn 2 lc rgb COLOR_BLUE
set style line RedBlackRelaxedBottomUp      dashtype 1 ps 1 lw 1 pt   8 pn 2 lc rgb COLOR_GREEN
set style line RedBlackRelaxedTopDown       dashtype 5 ps 1 lw 1 pt   9 pn 2 lc rgb COLOR_CYAN
set style line LBSTBottomUp                 dashtype 1 ps 1 lw 1 pt  10 pn 2 lc rgb COLOR_GREEN
set style line LBSTTopDown                  dashtype 5 ps 1 lw 1 pt  11 pn 2 lc rgb COLOR_BLUE
set style line LBSTRelaxed                  dashtype 1 ps 1 lw 1 pt  13 pn 2 lc rgb COLOR_RED
set style line TreapTopDown                 dashtype 5 ps 1 lw 1 pt  14 pn 2 lc rgb COLOR_RED
set style line TreapFingerTree              dashtype 1 ps 1 lw 1 pt  16 pn 2 lc rgb COLOR_BLUE
set style line Randomized                   dashtype 1 ps 1 lw 1 pt  17 pn 2 lc rgb COLOR_YELLOW
set style line Zip                          dashtype 1 ps 1 lw 1 pt  18 pn 2 lc rgb COLOR_GREEN
set style line Splay                        dashtype 1 ps 1 lw 1 pt  19 pn 2 lc rgb COLOR_PINK
set style line Conc                         dashtype 3 ps 1 lw 1 pt  20 pn 2 lc rgb COLOR_PURPLE


##################################################################
#
#           GROUPS
#
##################################################################

AVLRedBlack            = "AVLBottomUp RedBlackBottomUp RedBlackTopDown"
AVLWeak                = "AVLBottomUp AVLWeakBottomUp AVLWeakTopDown"
AVLRelaxed             = "AVLBottomUp AVLRelaxedBottomUp AVLRelaxedTopDown"
RedBlackRelaxed        = "RedBlackBottomUp RedBlackTopDown RedBlackRelaxedBottomUp RedBlackRelaxedTopDown"

RankBalanced          = "AVLWeakBottomUp AVLWeakTopDown AVLRelaxedBottomUp RedBlackRelaxedTopDown"
HeightBalanced        = "AVLBottomUp RedBlackBottomUp AVLWeakBottomUp AVLRelaxedBottomUp Conc"


WeightBalanced         = "LBSTBottomUp LBSTTopDown LBSTRelaxed"
WeightBalancedRelaxed  = "LBSTBottomUp LBSTRelaxed"
Probabilistic          = "TreapTopDown TreapFingerTree Randomized Zip"
Combination            = "AVLRelaxedBottomUp AVLWeakBottomUp LBSTBottomUp TreapTopDown"
CombinationSplay       = "AVLRelaxedBottomUp AVLWeakBottomUp LBSTBottomUp TreapTopDown Splay"
SizeOnly               = "LBSTBottomUp LBSTRelaxed Randomized Splay"

GROUPS = "AVLWeakRedBlack AVLRelaxed RedBlack"
GROUPS = "AVLRedBlack AVLWeak AVLRelaxed RedBlackRelaxed"

OPERATIONS = "Insert InsertPersistent InsertDelete InsertDeletePersistent InsertDeleteCycles InsertDeleteCyclesPersistent InsertDeleteSearch InsertDeleteSearchPersistent"
OPERATIONS = "Insert InsertPersistent InsertDelete InsertDeletePersistent InsertDeleteCycles InsertDeleteCyclesPersistent InsertDeleteSearch InsertDeleteSearchPersistent"

DISTRIBUTIONS = "Uniform Normal Skewed Zipf Maximum"
DISTRIBUTIONS = "Uniform Normal Skewed Zipf Maximum"

do for [GROUP in GROUPS] {

    do for [OPERATION in OPERATIONS] {

        ##################################################################
        #
        #           Allocations
        #
        ##################################################################

        MEASUREMENT = 'Allocations'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold Allocations } / log_2Size"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Allocations')/column('Step')/log2(column('Size'))"

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           MAXIMUM PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = 'MaximumPathLength'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold Maximum Path Length } / log_2Size"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('MaximumPathLength')/log2(column('Size'))"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           AVERAGE PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = 'AveragePathLength'

        set xlabel "Operations × 10^5"
        set ylabel "{/:Bold AveragePathLength } / log_2Size"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('AveragePathLength')/log2(column('Size'))"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           ROTATIONS
        #
        ##################################################################

        MEASUREMENT = "Rotations"

        set xlabel "{Operations × 10^5}"
        set ylabel "{/:Bold Rotations} / Operation"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Rotations')/column('Step')"

        DATA = sprintf("benchmarks/data/operations/measurements/%s", OPERATION)

        set format y2 "%.2f"

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        ##################################################################
        #
        #           DURATION
        #
        ##################################################################

        MEASUREMENT = 'Duration'

        set xlabel "Operations / 10^5"
        set ylabel "{/:Bold Duration } in milliseconds"

        x = "column('Position')/(column('Scale')/10)"
        y = "column('Duration')/1000000" # nano / 10^6 is 1k

        set format y2 "%.2f"

        DATA = sprintf("benchmarks/data/operations/benchmarks/%s", OPERATION)

        SMOOTH = "sbezier"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "unique"
        load "benchmarks/plot/operations/lines.gnuplot"

        SMOOTH = "cumulative"
        load "benchmarks/plot/operations/lines.gnuplot"
    }
}