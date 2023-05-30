#
#
#
#
#
#

array Partition[3] = [ "Median", "Height", "Weight" ]

array All[4] = [ "Median", "Height", "Weight", "DSW" ]

array Groups[2] = [ "Partition", "All" ]

array Distributions[5] = [ "Uniform", "Normal", "Skewed", "Zipf", "Maximum" ]

array Operations[1] = [ "Balance" ]



do for [Operation=1:|Operations|] {

    OPERATION = Operations[Operation]

    do for [Group=1:|Groups|] {

        GROUP = Groups[Group]

        ##################################################################
        #
        #           PARTITION COUNT
        #
        ##################################################################


        MEASUREMENT = "PartitionCount"

        set xlabel "{Size × 10^5}"
        set ylabel ""
        set title "{/:Bold Partition Count} / Size"

        set format y2 "%.2f"

        x = "(column('Size')/(column('Scale')/10))"
        y = "(column('PartitionCount')/column('Size'))"

        data_dir = "benchmarks/data/measurements"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_dots.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"


        ##################################################################
        #
        #           TOTAL PARTITION DEPTH
        #
        ##################################################################

        MEASUREMENT = "TotalPartitionDepth"

        set xlabel "{Size × 10^5}"
        set ylabel ""
        set title "{/:Bold Partition Depth} / Size"

        x = "(column('Size')/(column('Scale')/10))"
        y = "(column('PartitionDepth')/column('Size'))"

        set format y2 "%.2f"

        data_dir = "benchmarks/data/measurements"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_dots.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"


        ##################################################################
        #
        #           AVERAGE PARTITION DEPTH
        #
        ##################################################################

        MEASUREMENT = "AveragePartitionDepth"

        set xlabel "{Size × 10^5}"
        set ylabel ""
        set title "{/:Bold Partition Depth } / Partition Count"

        x = "(column('Size')/(column('Scale')/10))"
        y = "(column('PartitionDepth')/column('PartitionCount'))"

        set format y2 "%.2f"

        data_dir = "benchmarks/data/measurements"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_dots.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"


        ##################################################################
        #
        #           ROTATIONS
        #
        ##################################################################

        MEASUREMENT = "Rotations"

        set xlabel "{Size × 10^5}"
        set ylabel ""
        set title "{/:Bold Rotations}"

        x = "(column('Size')/(column('Scale')/10))"
        y = "(column('Rotations'))"

        set format y2 "%.0f"

        data_dir = "benchmarks/data/measurements"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_dots.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"


        ##################################################################
        #
        #           MAXIMUM PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = "MaximumPathLength"

        set xlabel "Size × 10^5"
        set ylabel ""
        set title "{/:Bold Maximum Path Length } / log_2Size"

        x = "(column('Size')/(column('Scale')/10))"
        y = "(column('MaximumPathLength')/log2(column('Size')))"

        set format y2 "%.2f"

        data_dir = "benchmarks/data/measurements"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_dots.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"

        ##################################################################
        #
        #           AVERAGE PATH LENGTH
        #
        ##################################################################

        MEASUREMENT = "AveragePathLength"

        set xlabel "Size × 10^5"
        set ylabel ""
        set title "{/:Bold Average Path Length } / log_2Size"

        x = "(column('Size')/(column('Scale')/10))"
        y = "(column('AveragePathLength')/log2(column('Size')))"

        set format y2 "%.2f"

        data_dir = "benchmarks/data/measurements"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_dots.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"


        ##################################################################
        #
        #           DURATION
        #
        ##################################################################

        MEASUREMENT = "Duration"

        set xlabel "Size × 10^5"
        set ylabel ""
        set title "{/:Bold Duration } / Size"

        x = "column('Size')/(column('Scale')/10)"
        y = "(column('Duration')/column('Size'))" # nanos to ms

        set format y2 "%.0f"

        data_dir = "benchmarks/data"

        smooth = "sbezier"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "unique"
        load "benchmarks/plot/balancers_lines.gnuplot"

        smooth = "cumulative"
        load "benchmarks/plot/balancers_lines.gnuplot"
    }
}