load "benchmarks/plot/balancers.gnuplot"

do for [Group=1:|Groups|] {

    GROUP = Groups[Group]

    ##################################################################
    #
    #           PARTITION COUNT
    #
    ##################################################################


    MEASUREMENT = "PartitionCount"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Partition Count} / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionCount')/column('Size')"

    set format y2 "%.2f"

    csv_dir = "benchmarks/csv/balancers/measurements"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"


    ##################################################################
    #
    #           TOTAL PARTITION DEPTH
    #
    ##################################################################

    MEASUREMENT = "TotalPartitionDepth"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Partition Depth} / Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionDepth')/column('Size')"

    set format y2 "%.2f"

    csv_dir = "benchmarks/csv/balancers/measurements"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"


    ##################################################################
    #
    #           AVERAGE PARTITION DEPTH
    #
    ##################################################################

    MEASUREMENT = "AveragePartitionDepth"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Partition Depth } / Partition Count"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('PartitionDepth')/column('PartitionCount')"

    set format y2 "%.2f"

    csv_dir = "benchmarks/csv/balancers/measurements"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"


    ##################################################################
    #
    #           ROTATIONS
    #
    ##################################################################

    MEASUREMENT = "Rotations"

    set xlabel "{Size × 10^4}"
    set ylabel ""

    set title "{/:Bold Rotations}"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('Rotations')"

    set format y2 "%.0f"

    csv_dir = "benchmarks/csv/balancers/measurements"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"


    ##################################################################
    #
    #           MAXIMUM PATH LENGTH
    #
    ##################################################################

    MEASUREMENT = "MaximumPathLength"

    set xlabel "Size × 10^4"
    set ylabel ""

    set title "{/:Bold Maximum Path Length } / log_2Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('MaximumPathLength')/log2(column('Size'))"

    set format y2 "%.2f"

    csv_dir = "benchmarks/csv/balancers/measurements"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"


    ##################################################################
    #
    #           AVERAGE PATH LENGTH
    #
    ##################################################################

    MEASUREMENT = "AveragePathLength"

    set xlabel "Size × 10^4"
    set ylabel ""

    set title "{/:Bold Average Path Length } / log_2Size"

    x = "column('Size')/(column('Scale')/10)"
    y = "column('AveragePathLength')/log2(column('Size'))"

    set format y2 "%.2f"

    csv_dir = "benchmarks/csv/balancers/measurements"

    smooth = "sbezier"
    load "benchmarks/plot/balancers/balancers_lines.gnuplot"

}
