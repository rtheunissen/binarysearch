Linear lists as balanced binary trees
===

Rudi Theunissen, June 2023

---

# Abstract

This project explores the design space[^1] of binary search trees to implement linear list data structures.
Binary search trees are still taught in computer science courses [https://sp23.datastructur.es] as if little has changed, but there are some algorithms that are more intuitive than the ones in the books.
Students often miss the opportunity to practice algorithm design because they are so distracted by seemingly arbitrary invariants and acronyms.
Almost certainly, anyone who knows about binary search trees can think of a time when they were not having fun learning about _red-black trees_.
Some new ideas have come along since then, but reference implementations are often difficult to find and therefore difficult to teach.
The goal of this project is to maintain a state-of-the-art algorithm reference that is easy to study and translate to other systems.

The balancing of binary search trees is an active, well-studied topic in computer science going back over 60 years.
Many algorithms to maintain balance have been invented over the years, but only a handful of usual suspects are usually found in textbooks and in practice.
Most students of computer science would be able to name one, maybe two or three of them, and almost certainly always the same ones.



# Table of Contents

1. Introduction
    - Memory
    - Linear lists
        - List operations
        - Dynamic arrays
        - Linked lists
    - Binary search trees
        - Relative position
        - Node structure
        - Traversal
        - Balance
        - Rotations
    - Persistence
        - Structural sharing
        - Immutability
        - Parent pointers
    - Concurrency
        - Contention
        - Recursion
    - Benchmarking
        - Access distributions
        - Operations
2. Strategies to restore balance
    - Balancing by rotation
        - Day-Stout-Warren
    - Balancing by recursive median partitioning
        - Height balance
        - Weight balance
3. Strategies to maintain balance
    - Join-based implementations
    - Rank-balanced trees
    - Weight-balanced trees
    - Randomly-balanced trees
    - Self-adjusting trees
    - Height-balanced trees
4. Evaluation
    5. Measurements
    6. Evaluation criteria
    5. Benchmarks
    7. Conclusions
       Source code
    9. Repository
    10. Animations
    11. Contributing
        -. Open problems
    - Missing proofs of correctness
    - Complexity of balancing by median partitioning
    - Similar analysis for set data structures
    - Translation to another language, perhaps Rust
    - Comparisons to other list data structures
      Notes
      References

                               PART 1

                             INTRODUCTION

## Memory

A computer program can use _memory_ to keep information while the program is running.
A program can allocate, access, and free individual pieces of memory as needed.
When some amount of memory is required, the program can ask an _allocator_ to reserve some memory and in return the allocator provides the program with an address with which to access that particular piece of memory. 
When a program is using a lot of memory, it suggests that it has allocated a lot more memory than it has freed and is holding on to a lot of information.

For example, when a web browser has many tabs open at the same time, all of that contextual information must be kept in memory and can only be freed when the user starts closing some tabs.
All the memory that was allocated by a program is usually freed when that program exits, but during the lifetime of the program there can be many requests to allocate and free pieces of memory. 

For the sake of illustration, we can think of memory as documents stored in a massive, indexed filing cabinet.
To allocate memory, we can ask a clerk known as "the allocator" to reserve space in the cabinet for us, enough for some known number of documents that we intend to work with. 
When we ask the allocator to reserve some space for us, assuming the cabinet is not full, we are provided with a number known as the _address_ which uniquely identifies that reserved space. 
When we are done with our work, we let the allocator know that the space we reserved is now free for someone else to use and whatever documents are still stored there can be discarded. 
The most important thing to be aware of is the capacity of the cabinet and the _cost_ of the allocator interacting with the memory. 

The "cost" here is simply energy; in the case of the cabinet clerk that would be kinetic energy, and in the case of a computer it would be electrical energy. 
Spending energy takes time, so we can assert that for a program to be fast it must minimize its interaction with memory. 
A program is considered _efficient_ when it meets its requirements using only a small amount of energy relative to the complexity of the problem it solves. 
Some problems can not avoid frequent interaction with memory, but an inefficient algorithm wastes energy by interacting with memory more often than it needs to.

> What does a program actually _do_ with memory? What is meant by "information" here exactly?

Consider a program to find the average age of all students in a school, where user data is stored in a file and every user has a date of birth as part of its data model. 
The program works by first loading all the students into memory, then inspects every user to determine their age based on the current date and the user's date of birth. 
Along the way, the program tracks the sum of each age and the number of students encountered.

At the end of the program, it prints the age sum divided by the number of students, then exits. 
The user data model also includes other information like name and email address. 
For the sake of illustration, we return to filing cabinet allocator and determine that we need one sheet of paper to record the information for one user. 
If there are N students in the school, we would need to allocate N sheets of paper to temporarily store all the user information in memory to calculate the average age. 
This algorithm is said to have _linear complexity_ because the required amount of time and space is linear in the size of the dataset: if the number of students were to double, so would the energy requirements of the algorithm. 

There is a more efficient approach, one not so eager in the way it reads the file where the user information is stored.
Instead of loading all the students into memory at the start of the program, the program scans the file one page at a time and only allocates enough working memory for one user.
As soon as it adds their age to the sum it can overwrite the user information with the next user from the file because it no longer needs that information. 
This algorithm requires only a _constant_ amount of memory but would still be linear in time because every student record _must_ be read. 
In this case, if the number of students were to double, the algorithm would take about twice as long to run but would not require any more memory than before and is therefore more efficient. 

As a final consideration to wrap up this concept, consider that a user might have an emergency contact which is also a user, and that user would then also have an emergency contact, and so forth. 
To avoid an infinitely described data model, the emergency contact of a user is stored not as a user but as an _address_ to a user: we call this a _reference_ or a _pointer_. 
When a program needs to know the name of a user's emergency contact it must first _dereference_ the memory address of that user's emergency contact to get the sheet of paper containing the user information of the emergency contact.

With all that in mind, we do not need to concern ourselves with the business of a particular program or the inner-workings of the memory allocator in the context of this project. 
The task at hand is to design data structures in memory that any program can use to organize information in a general way.

## Linear lists

This project explores specifically the implementation of a _list_ data type[^2]: a discrete, linear, ordered sequence.
Lists are found everywhere in computer programs and algorithm design.  
Using the example from the previous section, the first algorithm organized student records as a list in memory, and the user data file from that example likely also organized the information as a list of students.
As long as there is enough memory available to remember the sequence, there is no conceptual limit to the size of a list.

The following operations describe all the ways that a program can interact with a list, known as its programming _interface_. 
A data structure must support every list operation to qualify as an implementation of a list. 
As an exercise, consider how you might implement these operations using index cards to remember a sequence of words.

    - `New` creates an empty list.
    - `Size` returns the number of values in the sequence.
    - `Select` reads the value at a position in the list.
    - `Update` writes a value at a position in the list.
    - `Insert` adds a new value at a position in the list, increasing the position of all values that follow. 
    - `Delete` removes the value at a position in a list, decreasing the position of all nodes that follow.
    
### Dynamic arrays

Perhaps the most common computer memory implementation of lists in practice is the _dynamic array_[^3].
An array is a block of contiguous memory allocated all at once and indexed directly by offset from the address of the array structure itself.
The amount of memory to allocate for an array is the length of the array multiplied by the memory required for an individual value.

To illustrate this, we could use a sheet of grid paper to represent the allocated array where every square is a space in
which a value could be stored, perhaps in this case a single character. The number of squares on the sheet is the length 
of the array and the number of squares with values in them is the size of the list. Given that a list is a sequence of values, 
there would be no gaps between each position, so starting from a blank sheet as an empty list we insert one value at a time
starting in the top left corner working our way across the page and down as we start filling up the memory of the array.

Arrays are particularly efficient for some list operations, but less so for others. Lookup and update, both variations of _search_,
is very efficient because the memory address of a value in the list can be calculated as the address of the array plus the 
position of the value in question. Inserting a value at the end of the list is very efficient also: simply write into the next
empty space and increment the size. However, consider what happens when a value is to be inserted somewhere _within_ the sequence.
Since there are no empty spaces, a space must first be created by moving every value from that position onwards one position towards the
end of the list, which is _linear_ in the size of the list. Delete is similarly inefficient, as well as split and join.

What makes an array _dynamic_ is the ability to increase its memory capacity by reallocating the entire sequence into a larger
block of memory, copying every value from the original array to the new allocation, followed by a free of the previous allocation.
Doing so is _linear_ in the size of the list but occurs infrequently enough to not consider it a major concern. Common ways to
mitigate this cost is to predict accurately an upper bound of the required length, or otherwise to double the previous length
whenever capacity runs out.  Insertion, expected to be efficient at the end of the list, is said to be _amortized_[^4] because the 
many low-cost operations make up for the infrequent large cost of reallocation. This characteristic of dynamic arrays is 
often used to teach amortized analysis.

### Linked lists

Found somewhere in every textbook on algorithms and data structures is the _linked list_[^5], consisting of _nodes_ 
that each contain a value and a pointer to the next node in the sequence. The data structure itself keeps a pointer
to the first node of the sequence, here referred to as the "root" node, as well as a counter of the total number 
of linked nodes in its sequence to support the _size_ operation. The last node of the list points to nothing because there is no next
value after it in the sequence. A pointer to nothing is a _null_ pointer or _nil_ in some languages. 
To insert a new value at a position within the sequence, start at the root and follow pointers one node at a time 
until the number of nodes encountered equals the target position. At this point, we can insert a new value by allocating 
a new node for it, pointing the current node to that node, and the new node to the previous next node of the current node. 
Deleting a node is similar: adjust the link from the node pointing to the node to be deleted to the next node after it,
thereby removing that node from the sequence of nodes reachable from the root.

Many of the operations of a linked list require that we follow a number of links equal to the target position and is
therefore _linear_ in the size of the list: the number of memory interactions grows at the same rate as the size of the list. 
The maximum path length is equal to the size of the list.
There is no way around this because the only way to reach a specific node is to follow the path that leads to it, 
starting from the root node since that is the only reference we have in hand at the start of the operation. 
There is no need to allocate a large amount of memory all at once or to reallocate anything because memory is allocated 
individually for each node as needed. The trade-off here is that this requires more allocations overall and often results
in memory fragmentation: the next node in the sequence might be physically far away from the current node, compared to an 
array where the next value is always immediately adjacent because the entire sequence is one contiguous block of memory.

To illustrate this, imagine a very long line of people in a field where everyone is pointing to the next person in line after them.
No one knows their current position other than the person at the very front and maybe the next few people after them, but 
only so because they can see and count the number of people before them. To find the person at some given position, 
someone (the program) would need to count people one at a time from the start of the line, each time following along to 
where that person is pointing. "What the point of all this pointing is exactly?". Keep in mind that the people in line might 
not all be standing neatly in order. In fact, the next person in line might be all the way on the other side of the field
and the program would need to walk very far to reach them, which requires a lot of energy and therefore takes a long time.

Imagine now that a ticketing agent is trying to give everyone some chance of being close to the front of the stage.
Whenever someone arrives to line up they are assigned a random number between zero and the number of people in line, 
indicating the number of people they are allowed to skip. Walking up to a random person in line would not help because
they would not know what their position is exactly, since others may have joined the line ahead of them since they got there.
The only way to operate this strategy is to start at the front of the line and count links along the way. 
Linked lists are simple to understand and program, but their linear complexity and memory cost often makes them 
non-viable in practice. 

Instead of starting the search at the front of the list, what if we started somewhere in the middle? 
Nodes to the right of this median would point to the next node in the sequence as they did before, 
but nodes to the left of the median would now point in the other direction towards the front of the list.
Doing so reduces the maximum path length because the program would now need to walk at most about halfway.
To achieve this, notice that the median node now has _two_ pointers: one going left and another going right.

## Binary search trees

Applying the same strategy _recursively_ to the left and right sides of the median node produces 
a structure known as a _binary search tree_, to which we finally arrive. The initial median node
is at the very top of the tree, becoming the _root_ node. Every node now has two links, 
one to its left subtree and another to its right subtree. 

### Node structure

```go
type Node struct {
   l *Node
   r *Node
   s uint64
   x any
}

00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 LEFT   | 
00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 RIGHT  | STRUCTURE
00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 SIZE   |
00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 DATA

```


### Traversal

The fundamental rule that every node must follow is that the value of the left node exists somewhere _earlier_ in the sequence, and the value of the right node exists somewhere _after_. 
When this rule is valid at every node in the tree, we get a total linear ordering of all values even though the paths branch in two directions every time.
The same sequence of the original linked list can therefore still be derived from the tree structure, making use of an _inorder tree traversal_.

Starting at the root, perform a traversal of the left subtree, then produce the root, then perform a traversal of the right subtree.
The first node produced by this algorithm is therefore the left-most node of the tree and the last node is the right-most node.
To illustrate this, we can draw a binary search tree on a grid where the horizontal coordinate is the position of the node in the sequence, and the vertical coordinate is the length of the longest path.
Dragging a vertical line across the page from left to right will intersect the nodes in sequential order. 
The only address we have in hand at the start of the traversal is the root, so to get to the left-most node we need to follow along the _left spine_ of the tree first, eventually coming back up to the root before descending to the left-most node of the right subtree, and so forth.

### Logarithmic path length

What is the maximum path length of this tree structure? 
Every time we partitioned a list around its median we halved its maximum path length, and then we halved the maximum path length of each half, and so on until we reached a single node that could not be partitioned any further.
The question becomes "how many times can we divide a number by 2 until we get to 1?".
This is known as the _binary logarithm_[https://en.wikipedia.org/wiki/Binary_logarithm], written as log₂ and sometimes referred to as "log with base 2". 
If a pizza has 8 slices, we can halve it once to get 4 slices, again to get 2 slices, and one more time to get 1 slice, and therefore the binary log of 8 is 3 because we could split the pizza in half three times. 
The binary log function is frequently encountered in computer science and information theory because it relates so closely to binary numbers in the base 2 numeral system; numbers made up of only 1's and 0's knows as "bits".
For example, the number of bits required to encode an integer `n` in binary is ⌊log₂(n)⌋ + 1.
The same reasoning applies to decimal numbers: the number of digits required to represent an integer is the floor of the log with base 10, similarly the number of times we can divide a number by 10 before we get to 1. 
When we divide a decimal number by 10, the representation shortens by 1 digit. 
Similarly, when we divide a binary number by 2, the representation shortens by 1 bit. 
A two-way branch at every node recursively halves the maximum path length, which results in a maximum path length _logarithmic_ in the size of the list: 
if the size of the list doubled, the search cost increases by one. More generally, the search cost grows _much_ slower than the size of the list.

Some nodes may not have a left or right subtree, suggesting that either or both links could be null, 
but the pointers themselves are still there because they are part of the allocated node structure. 
A node that has neither a left or right subtree is called a _leaf_ node.

### Inserting a new node into a binary tree

Recall that inserting a value into in a dynamic array requires the program to first shift all values to the
right of that position to create space for the new value. The null pointers of a binary tree is exactly these gaps,
where number of null pointers in a binary tree is always one more than the number of nodes. 
There is a gap at the front of the sequence, in-between every value, and a gap at the end of the sequence.
Inserting a node into a binary tree replaces one of these null pointers with a new leaf node,
thereby occupying one of the previously existing gaps but also creating two more gaps.
To illustrate this, we can draw the missing pointers on any tree that does not include them throughout
this project, count them, and observe that there is always one more than the number of nodes in the tree.

### Relative position

We now have a data structure with nodes pointing to other nodes, but it is not yet clear how this operates as a list.
Traversing a linked list in search of a position is simple because every time we follow a link we count exactly
+1 node. To implement search by position on a binary search tree, we need to be able to count towards a position. 
When we follow a link in either direction along the search path, we skip all the nodes of the opposite branch. 
At the root of the tree, the first branch skips about half the nodes of the entire sequence, then the second branch skips 
about a quarter, then an eight, all the way down the tree until we find the node at the position we are searching for. 

This is known as _binary search_, which is a well-known algorithm completely independent of linked nodes and tree structures.
Many of us have applied this algorithm in our every-day lives perhaps without realizing it. Consider an old-school dictionary
or telephone directory where all the records are printed in alphabetical order. How would you search for a particular entry?
If we started on the first page and searched one page at a time it would be a _linear search_, taking time proportional
to the size of the directory. Our intuition might suggest to us that there is a smarter way to go about this, that we can achieve the
same result by doing less work: start somewhere in the middle of the book, perhaps not the exact median page but somewhere close to it,
then determine whether the page we are looking for is before or after that page. Because the directory is _ordered_, 
we know that we can reject either the left or right half entirely. Repeating this step on the resulting half once again 
divides the search space again in half, until eventually we close in on the page we are looking for.
To implement binary search, we need to determine whether to branch to the left or the right; do we reject the left half or the right half?

In the context of a list implementation, this determination requires some _comparison by position_, for which we do not yet have the required information.
In the case of a linear search, we tracked the current search position by counting nodes one at a time, but a binary search tree can skip _multiple_ nodes at a time.
To apply a similar counting strategy we need to know how many nodes are in the left subtree: if we were to reject all the nodes of the left subtree, would we skip too far ahead in the sequence?

We can follow this algorithm to search by position:

      If the position we are searching for is equal to the number of nodes in the left subtree, we have found the node we were looking for because skipping all the nodes of the left subtree would skip ahead to exactly this position in the sequence.
      Otherwise, if the position is greater than the size of the left subtree, we know that we can reject the entire left subtree because even if we skipped all those nodes we would still need to skip further ahead in the sequence, and therefore the node we are looking for must be somewhere in the right subtree. 
      In this case, reduce the search position by one plus the size of the left subtree and follow the link to the right to continue the search. 
      Otherwise, the position is still less than the size of the left subtree, suggesting that we need to seek towards the front of the list because our current position is still too far ahead in the sequence. 
      In this case, follow the link to the left and continue the search. 

This brings us to the root of our first problem: how can we know the relative position of the current node?
We previously structured a binary tree from an existing linked list by choosing the exact median at every step, but binary search does not have such a strict requirement to be effective.
Even a somewhat poor approximation of the median might only increase the search cost by a few steps.
Therefore, we must accept that the current node along the search path might not be the exact median of its subtree, and therefore we can not know its position unless we record that information somewhere.

Perhaps the most common approach to solve this is to store in every node the size of that subtree. [https://en.wikipedia.org/wiki/Order_statistic_tree] [340 Chapter 14 Augmenting Data Structures CLRS].
Inserting a new node requires that we then increment the size field of every node along the search path because the new node that will eventually be 
attached at the bottom of the tree would be a common descendant, thereby increasing each of their size by 1.
To determine the relative position of a node we can dereference its left link to read the size field of that node, wherein lies a fundamental weakness: 
to know the size of the left subtree we must first dereference it, even though the search path might end up branching to the right. 
This is a weakness because we have to spend energy to look up the address of a node to know its size.

We could instead store in each node the size of its left subtree specifically, as suggested by Knuth in [] and Crane in [].
Keeping a separate _size_ field in the tree alongside the reference to the root node allows us to track the current subtree size at each step along the search path. 
The size of the right subtree can be calculated as needed as the size of the subtree minus the known size of the left subtree, minus one.
This approach allows us to know the subtree sizes of both the left and right nodes without the need to dereference either of them.
Insertion then only increments the size field of a node when branching to the left because inserting to the right would 
not affect the size of the left subtree. This approach therefore reduces memory interaction in exchange for a slightly 
more complex tree structure and some inherent asymmetry - a good trade.

A third design candidate is to store in each node its position relative its parent, where the parent is the node that points to it.
This representation results in a left node having a negative position equal to the negative size of its right subtree minus one, and a right node having a positive position equal to the size of its left subtree plus one.
Following a convention where the root node of a tree always has a positive position, the absolute position of any node is then equal to the sum of all relative positions along its path from the root.
This strategy is symmetrical, intuitive, and provides one bonus insight: a node is known to be a left or right descendant based on the sign of its position. 
However, there are two downsides to this representation: the first is that we require one bit to store the sign of the position, thereby halving the utilization of the integer field, 
and the second is that the resulting algorithms require in some cases additional checks and arguments to indicate and normalize node orientation.
Insertion using this strategy would increment the size field when a search path branches left at a right node, and symmetrically decrement the size when branching right at a left node.
For example, inserting a node at the very front of the sequence would increment the size of the root node because the size of its left subtree is increasing, then descend along the left spine of the tree without the need to increment any others.
Similarly, inserting a node at the very end of the sequence would not increment the size of any node because all the nodes along the right spine including the root have positive positions indicating the size of their left subtrees, unaffected by an insertion to the right.

This last approach is rare but not unheard of.
In 2004, Jörg Schmücker proposed a list implementation using this exact approach to the Apache Commons Collections library in 2004 [https://markmail.org/message/43ux2i3rbsigtotu?q=TreeList+list:org%2Eapache%2Ecommons%2Edev/&page=4#query:TreeList%20list%3Aorg.apache.commons.dev%2F+page:4+mid:mv2nw4ajw2kywmku+state:results],
which is still part of the main branch at the time of this writing [https://github.com/apache/commons-collections/blob/3a5c5c2838d0dacbed2722c4f860d36d0c32f325/src/main/java/org/apache/commons/collections4/list/TreeList.java].
In 1997, in section 6.3 of [Randomized Binary Search Trees], Martinez and Roura describe the use of an "orientation bit" in every node to indicate which of the two subtrees the size field is referring to.
They suggest to flip the orientation at every step along the search path as needed to always store the size of the subtree _not_ increasing in size, thereby leaving behind valid size information in the wake of a failure.

This project uses the approach where every node consistently stores the size of its left subtree:

    - No need to dereference a node to know the size of its subtree.
    - No need to consider whether a node is a left or right link.
    - No need to adjust the size field of a node when branching to the right.

The following illustration shows a binary search tree where every node is annotated with the size of its left subtree,
followed by a function that implements the search algorithm described above.

```go
// `p` is the current node along the search path.
// `i` is the current search position, also the distance from the result.   
func lookup(p *Node, i uint64) *Node {
   for {
      if i == p.s {
         return p           
      }
      if i < p.s {           
         p = p.l            
      } else {               
         i = i - p.s - 1    
         p = p.r
      }
   }
}
```

### Complexity in comparison

TODO A table showing linked list vs dynamic array vs binary search tree complexities as: CONSTANT, LINEAR, or LOGARITHMIC

## Persistence 

A _persistent data structure_ can create many independent versions of itself over time, where changes in a future version would not be visible to a program still referencing an older version.
Persistence adds the dimension of time as a history, where any version can still be modified to create new versions from that point in time without affecting other versions.
A program can preserve the current state by first allocating a duplicate structure before making changes, thereby producing a new version entirely.
For this to be efficient, a data structure must be _cheap to copy_ and implicitly share memory between versions over time. 


https://en.wikipedia.org/wiki/Persistent_data_structure

### Reference counting

To avoid copying all the nodes when copying a tree, we allow trees to share common subtrees in memory over time.
Sme tree far in the future might still point to a node that was allocated in a much earlier version.
We need to keep track of these references, so we store in every node a _reference count_ indicating the number of other trees that also reference that node.
When the reference count of a node is zero, it suggests that no other trees are aware of that node, so the program can modify it without the need to copy it first.
No copying will occur and all modifications will be in-place if a tree never shares a node with another tree.
When the reference count of a node is greater than zero, it suggests that another tree has a dependency on it.
Before making a change to this node, which at this point would also change the other trees that depend on it, we must first detach it from its history.
This can be done by (1) replacing the node by a duplicate of its node structure, (2) incrementing the reference counts of the left and right links, (3) decrementing the reference count of the original node, and (4) setting the reference count of the copy to zero.
There is now one fewer tree depending on that node specifically, and because the reference count of the new node is zero, the program can make changes to it as needed.

https://en.wikipedia.org/wiki/Copy-on-write
https://en.wikipedia.org/wiki/Reference_counting
https://en.wikipedia.org/wiki/Immutable_object

### Path copying

Consider now that some node is being modified in the depths of some version of a tree, but that node is shared, so must first be copied.
How can the root node of the tree know about that new branch? 
How could it reach it? 
Looking at it from the other side, a node replaces another node on the search path when it is copied, so there is some parent node now pointing to that copy.
The program would need to change the left link of that parent node to point to the copy, so it too must first be copied.
This continues until it cascades all the way up to the root of the tree. 
There is only one rule to consider when thinking about path copying: during an operation, all paths that lead to a modification must be copied.
Every node has a unique path from the root, so for any operation we can mark the nodes that would need to be modified and trace their paths back up to the root; 
it is these paths that would need to be copied so that they all belong to the same, new version of the tree that includes those modifications. 

https://en.wikipedia.org/wiki/Persistent_data_structure#:~:text=in%20the%20array.-,Path%20copying,-%5Bedit%5D

### Parent pointers

Many implementations of binary search tree algorithms involve one additional pointer: the _parent_ pointer, pointing back up to the node that points to it.
Path copying requires that _during an operation, all paths that lead to a modification must be copied_, but with parent pointers all paths lead to all nodes.
Parent pointers create cycles between nodes, which is not compatible with path-copying and therefore not part of the node structure used in this project.
Avoiding parent pointers also saves on memory interaction because there are fewer pointer relationships to maintain, and no need to allocate space for the pointer in the node structure. 

## Concurrency

### Contention

A data structure that supports _concurrency_ provides in some way the ability to run multiple operations at the same time.
This is not exactly _parallelism_, which is where the work of a single operation is divided up and worked on concurrently, together towards the same outcome.
Concurrency, for example, is where an operation to insert a new value does not prevent another similar operation from starting before the previous one ends.
To illustrate this idea on a binary search tree, imagine two insert operations starting at the same time.
One of these operations must win the race to the root node, and the other operation must wait in line right behind it: this is called _contention_.
The first insert operation must branch to the left of the root, so it increments the size of the left subtree and follows the left link, never to be seen again.
The second operation is clear to operate on the root node as soon as the first operation follows that left link: 
if it too must branch left then it will likely be blocked by the first operation again, but if the second operation branches to the right then these two operations are forever independent.
An operation could modify a node in one part of the tree while another operation is modifying a different part of the same tree.

### Recursion

Binary tree algorithms often use _recursion_ to implement operations as two phases: (1) a search phase downward from the root, and (2) a maintenance phase upward towards the root.
A blocking operation would need to wait for the algorithm to go all the way down, then come all the way back up to make final adjustments.
For this reason, we choose whenever possible, to implement algorithms as one iterative top-down phase, combining the search and maintenance phases into one. 
Recursive implementations are still valuable because they help us to learn the algorithms and verify the iterative solutions.

---



                                            

                                  PART 2

                      STRATEGIES TO RESTORE BALANCE


Consider then what would happen to the structure of the tree if we repeatedly inserted many nodes at the end of the sequence.
Eventually, the tree structure becomes _unbalanced_, where too many branches are too far from the median of their sequence.
Starting from an empty tree, always inserting at the end of the sequence would create precisely a linked list.
Over time, even a perfectly balanced tree might become unbalanced when nodes are inserted, deleted, split, and joined.

The same tree can be arranged in many unique ways[Catalan numbers] without changing the linear order of its nodes.
There is always a way to arrange a tree so that any of its nodes could be the root.
For example, using the first node of a sequence as the root would result in a tree with no left subtree.

 
                  a  b  c         a  b  c           a  b  c
```
                 (a)                (b)                  (c)
                   \               /   \                /
                    (b)          (a)   (c)            (b)
                      \                              /
                       (c)                         (a)
```

Keep in mind that a node only keeps pointers to other nodes, not the nodes themselves, so to get the information of another node we need first retrieve it from memory.
Every time the program follows a link from one node to another, it must interact with memory to dereference that next node.
Generally, a tree search uses less energy when there are fewer links to follow per search. 
Balance is therefore a measure of low _average path length_: on average, how many links would a program need to follow?
The path length of a node is the number of links to reach it from the root.
The average path length is the sum of all path lengths divided by the size of the tree.

To restore balance, we need a strategy to reduce the average path length by adjusting the branches of a tree without changing its underlying sequence.
Some balancing strategies spend a lot of energy to achieve perfect balance, while others spend less energy by being more relaxed.
We can evaluate these strategies as a trade-off between balance and energy: to what extent can balance be relaxed before search cost becomes too much?

### Rotations

A tree _rotation_ is a minor local branch adjustment that changes the path lengths of some nodes without changing their order.
Consider the three of nodes in 1 and the transformation required to form the tree in 2: 
move (c) up from the right, pushing (a) down to the left, dragging (d) along with it, and (b) moves across.
The transformation from 2 to 3 is similar:
move (d) up from the right, pushing (c) and its left subtree (a) down to the left.
Rotating back in the other direction at (d) and then (c) revert back to 1.
At all stages, the sequence was (a,b,c,d).

```

          1                      2                     3
                                              
1         (a)                     (c)                    (d)
             \                   /   \                   /
2            (c)               (a)   (d)               (c)       
             /  \                \                     /       
3          (b)  (d)               (b)                (a)     
                                                      \       
                                                      (b)    
                                                      
                                                      
```


A well-known algorithm to restore balance using tree rotations is the _Day-Stout-Warren_ algorithm or simply DSW, designed by Quentin F. Stout and Bette Warren in 1986, [1] based on work done by Colin Day in 1976.
This algorithm first transforms the tree into a linked-list, then transforms that linked-list into a balanced tree, all using rotations.                                         

### Partitioning

In 1980, Stephenson [] presented an algorithm to always insert a new node at the root of the tree by splitting the tree in two: nodes that occur before the new node and nodes that occur after, then setting those two trees as the left and right subtrees of the new node. 
A variation of this algorithm is called _partition_, which moves the node at a given position to the root of its tree in a single top-down pass:

      Allocate two nodes, L and R.

      Start a binary search with the root of the tree as P.

      When the search position is equal to the position of P:
         The node to become the root has been found. 
         Append the left subtree of P to the right of L. 
         Append the right subtree of P to the left of R. 
         Set the left link of P to L.
         Set the right link of P to R.
         Set the position of P to the initial search position.
         Done.

      When the search branches left:
         Append P to left of R.
         Reduce the position of P by the search position + 1.
         Follow P left.

      When the search branches right:
         Append P to the right of L.
         Reduce the search position by the position of P + 1.
         Follow P right.


```
func partition(p *Node, i uint64) *Node {
   n := Node{s: i}
   l := &n
   r := &n
   for i != p.s {
      if i < p.s {
         p.s = p.s - i - 1
         r.l = p
         r = r.l
         p = p.l
      } else {
         i = i - p.s - 1
         l.r = p
         l = l.r
         p = p.r
      }
   }
   r.l = p.r
   l.r = p.l
   p.l = n.r
   p.r = n.l
   p.s = n.s
   return p
}
```


### Median balance

Definition:
   A node is **median-balanced** if the _size_ of its left and right subtrees differ by no more than 1.

Definition:
   A tree is median-balanced if all of its nodes are median-balanced.

More recently in 2017, Ivo Muusse published an algorithm for balancing a binary search tree [] that uses _partition_ to replace every node by the median of its subtree.
Consider a node of size 100 at position 10, suggesting a left subtree size of 10.
The median node is at position 50, so the algorithm moves that node to the top to replace the original node, which is now somewhere in the left subtree.
The number of nodes is now evenly distributed between the left and right subtrees, but there might be nodes within those subtrees that are not balanced.
Repeating the same steps recursively in the left and right subtrees results in a median-balanced tree.

         Start with the root of the tree as P.
         Partition the median node of P if P is not balanced.
         Balance the left subtree of P.
         Balance the right subtree of P.

This algorithm has multiple useful properties:
(1) it is general for any definition of balance based on subtree size;
(2) it works well on trees that are already somewhat balanced;
(3) works strictly top-down which is great for concurrency because the result of a partition is a valid node and the algorithm only descends from there;
(4) subtree balancing is independent so could be done in parallel;
(5) the balancing operation can be cancelled without invalidating the tree.

A median-balanced tree is perfectly balanced because there is no arrangement of that tree with a lower average path length.
However, there are some arrangements that have the same average path length but are not strictly median-balanced.

### Height balance

Definition:
   The **height** of a node is equal to its maximum path length.

Definition:
   The height of a tree is the height of its root.

Definition:
   A node is **height-balanced** if the height of its left and right subtrees differ by no more than 1.

Definition:
   A tree is height-balanced if all of its nodes are height-balanced.


```
                  1                            2



                 (b)                          (c)       
                /   \                        /   \      
              (a)   (d)                    (b)   (d)    
                   /   \                   /       \   
                 (c)   (e)               (a)       (e) 
                                                
            
            HEIGHT-BALANCED             MEDIAN-BALANCED
            
```

Consider the trees in [Figure] that both form the same sequence: (a,b,c,d,e).
Both trees have a size of 5, average path length of 6/5, and a maximum path length of 2.
However, the first tree is not median-balanced because 5/2 is 2 so the median is (c), but the root is (b).
The median-balancing strategy would partition at (b) to make (c) the root, but the average path length stays the same.
This partitioning step is therefore work that is not productive because balance did not improve.
We can continue to use the same partition-based balancing algorithm, but the definition of balance must change to only partition if the node is not already height-balanced.
Any node that is _not_ height-balanced is partitioned to become median-balanced, which results in a height-balanced tree.
The problem to solve is then <mark>to determine whether a node is height-balanced using only the size of its left and right subtrees</mark>.

Muusse solves this in []: pretend that both subtrees are already strictly height-balanced, then compare their minimum and maximum heights.
When a tree is height-balanced, all the layers are of the tree are full except for maybe the bottom level, where some gaps might exist.
The bottom level of the larger subtree gets completed with _ceil_, and the bottom level of the smaller subtree gets emptied with _floor_.
By comparing these heights we can determine if the height difference is greater than 1:

      HeightBalanced(x,y) := ⌈log₂(max(x,y)+1)⌉ - ⌊log₂(min(x,y)+1)⌋ ≤ 1

Without loss of generality, assume that x > y:

      HeightBalanced(x,y) := ⌈log₂(x+1)⌉ - ⌊log₂(y+1)⌋ ≤ 1

This function, as presented above and by Muusse in [], can be simplified using an identity of log₂ where `⌈log₂(x+1)⌉ ≡ ⌊log₂(x)⌋+1` [].
The result is the same whether you complete the bottom level with _ceil_, or empty the bottom level with _floor_ and add 1. 

      HeightBalanced(x,y) := ⌊log₂(x)⌋ - ⌊log₂(y+1)⌋ ≤ 0

      HeightBalanced(x,y) := ⌊log₂(x)⌋ ≤ ⌊log₂(y+1)⌋

The calculation is now whether the floor of the log₂ of `x` is less than or equal to the floor of the log₂ of `y+1`.
Given that the number of bits required to encode an integer `i` in binary is `⌊log₂(i)⌋+1`, we can add `+1` to both sides of the inequality to see that what we are comparing is the number of bits required to encode `x` and `y+1`.
The most-significant bit, or the MSB, is the left-most bit set to 1, starting from position 1 on the right counting left.
All the bits to the left of the MSB will be 0, so the position of the MSB is equal to the number of bits required, and therefore equal to the floor of its log₂.

For example:

      00001101 = 13, because 8 + 4 + 0 + 1 = 13
          ↑
         MSB

            The MSB of 13 is at position 4
            log₂(13) = ~3.7, so ⌊log₂(13)⌋+1 = 4
            The number of bits required to encode the number 13 in binary is 4


      00010001 = 17, because 16 + 0 + 0 + 0 + 1 = 17
         ↑ 
        MSB

            The MSB of 17 is at position 5
            log₂(17) = ~4.1, so ⌊log₂(17)⌋+1 = 5
            The number of bits required to encode the number 17 in binary is 5


Using this information, we can determine whether a node is height-balanced by comparing the MSB of the size of each subtree.
More generally, we need to determine if the MSB of one integer is less than or equal to the MSB of another.

      HeightBalanced(x,y) := ⌊log₂(x)⌋ ≤ ⌊log₂(y+1)⌋

      HeightBalanced(x,y) := MSB(x) ≤ MSB(y+1)

This allows us to determine height-balance without the need to actually calculate either logarithm, which would require slow floating-point operations.
There are a few ways to compare the MSB of the integers using bitwise operations [][][]:

      SmallerMSB(x, y) := x < y && ((x & y) << 1) < y    Roura, S. (2001). A New Method for Balancing Binary Search Trees. ICALP.
      SmallerMSB(x, y) := x < y && x < (x ^ y)           Chan, T.M. (2002). Closest-point problems simplified on the RAM. SODA '02.
      SmallerMSB(x, y) := x < (~x & y)                   Warren, H.S. (2002). Hacker's Delight. 2nd Edition, section 5-3.
   
Using the third approach involves only one inequality and benchmarks slightly faster than the others.
Putting it all together, we can use the following expression to determine if a node with two height-balanced subtrees is height-balanced:

      HeightBalanced(x, y) := (x & ~(y+1)) ≤ (y+1)


### Weight balance

Another consideration is to keep the size of each subtree within a constant factor of the other.
For example, we might consider a node to be balanced if the size of one subtree is no more than twice that of the other.
This strategy is commonly known as _weight balancing_, originally invented as "trees of bounded balance" in 1980 [] ??
In 2001, Roura presented a variant of weight balance that directly considers the logarithm of the subtree size.

// Steps, stick figures?

Definition:
   A node is **weight-balanced** if the bit positions of the MSB of each subtree size differ by no more than 1.

This definition of balance is more relaxed but still maintains a height upper-bound of `2×⌊log₂(size)⌋`.
For example, a node with subtree sizes 17 and 8 is balanced, but sizes 17 and 7 are not balanced because the most significant bits are too far apart.


                     ↓↓                    ↓ ↓
                  00010001 = 17         00010001 = 17
                  00001000 = 8          00000111 = 7

                  BALANCED              UNBALANCED 


Without loss of generality, assume that x > y.
The MSB of `x` is then either at the same position as the MSB of `y` or further left.
Shifting all the bits of `x` one step to the right then either aligns the MSB of `x` with the MSB of `y`, or moves the MSB of `x` one step to the right of the MSB of `y` if they were already aligned.
After the shift, if the MSB of `x` is either aligned with the MSB of `y` or further to the right then the sizes are balanced.
Otherwise, the MSB of `x` is still to the left of the MSB of `y`, suggesting that it was too far away and therefore unbalanced.
Balance is therefore met if the MSB of `x` shifted to the right by one step is less than or equal to the MSB of `y`. 

      WeightBalanced(x, y) := MSB(x >> 1) ≤ MSB(y)

      WeightBalanced(x, y) := ((x >> 1) & ~y) ≤ y  


### Analysis

We now have three definitions of balance for a partition-based balancing algorithm: `median`, `height`, and `weight`.
For measurements and analysis, trees are created in size increments of 100 all the way up to 1,000,000 for a total of 10,000 samples.
At each increment, a random tree is balanced independently by each definition, then measured to capture: (1) partition count, (2) partition depth, (3) path length, and (4) time taken.

TODO: Mention DSW

[image of balancer graphs]

Observations on balancing random trees by partition:

   (1) Balancing by height results in the same height and average path length as by median.
   (2) Balancing by height partitions more often but total partition depth is lower than by median.
   (3) Balancing by height is slower than balancing by median. **Unexpected!**

   (2) Balancing by weight results in slightly higher height and average path length.
   (5) Balancing by median is faster than balancing by height. 
   (6) Balancing by weight is faster than balancing by height and median. 
   (7) Complexity appears to be linear.

The expectation was that height-balance uses fewer partitioning steps to achieve the same path length as median-balance, but this does not appear to be the case.
Height-balancing interacts with fewer nodes during each partition on average, but partitions more often enough to be less efficient than median-balancing overall.





                                  PART 3

               STRATEGIES TO MAINTAIN BALANCE OVER TIME


Restoring balance to an existing tree is useful but expensive — the entire tree must be traversed.
There is also the question of _when_ exactly to perform such a balancing operation, to avoid balancing too often or not often enough.
Consider instead of either balancing the entire tree or not at all, to balance incrementally and thereby distribute the cost of balancing.
During an update, an algorithm can make structural changes along the path to maintain balance over time.
Some strategies make structural changes all over the place, while others are more selective.
A valid balancing algorithm guarantees that the tree is always left in a balanced state.
A program can then _know_ that a given tree is balanced because it is always balanced.

This part of the project introduces a selection of strategies to compare in practice.

https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree

### Height-balanced trees

Since the invention of AVL trees in 1962, 

Perhaps the most common strategies to maintain binary tree balance over time are AVL trees [], and red-black trees [].
AVL trees store the height in every node and uses rotations along the search path to fix balance as needed.
Red-black trees store a color in every node (red or black), and enforces rules around these colors to determine where to rotate along the search path to fix balance as needed.
There are many debates about which is better, where the resolution is usually that _it depends_.
Most implementations in practice seem to use red-black trees [], [], [].

In 2015, Haeupler, Sen, and Tarjan published a unifying framework for a class of balanced trees called _rank-balanced trees_ [].
Rank-balanced trees store in every node an integer _rank_, then uses rank difference rules between nodes to define height-based balance.
Different invariants on rank differences are equivalent to AVL trees, red-black trees, and other kinds of balanced trees. 
By relaxing AVL trees, they obtain a new kind of balanced binary tree: a _weak AVL_, or WAVL tree, which combines the best properties of AVL and red-black trees into a single, unified algorithm.
The height bound of a WAVL tree degrades gracefully from that of an AVL tree as the number of deletions increases, but is never worse than that of a red-black tree. 
Both insertion and deletion in WAVL trees can be implemented purely top-down and in-place, without recursion or a stack otherwise.

In 2016, Sen, Tarjan, and Kim explored rank-balanced trees further to avoid rebalancing on deletion [].
These variants are called _relaxed rank-balanced_ trees because they only restore balance after insertions.
The height bounds are therefore not proportional to the size of the trees, but the total number of insertions instead.
Algorithms for _relaxed AVL_ or RAVL trees are described, as well as those for _relaxed red-black_ trees.

We implement the following variants and evaluate their performance and behavior:
   
      AVL               : bottom-up
      Weak AVL          : bottom-up, top-down, join-based
      Relaxed AVL       : bottom-up, top-down
      Relaxed RedBlack  : bottom-up, top-down
    
[ Benchmarks of rank-balanced trees]



### Weight-balanced trees

### Random-balanced trees

### Self-adjusting trees

### Join-based balancing

### External search trees


























[^2]: https://en.wikipedia.org/wiki/List_(abstract_data_type)

[^3]: https://en.wikipedia.org/wiki/Dynamic_array

[^4]: https://en.wikipedia.org/wiki/Amortized_analysis

[^5]: https://en.wikipedia.org/wiki/Linked_list


https://en.wikipedia.org/wiki/Copy-on-write

https://en.wikipedia.org/wiki/Order_statistic_tree

"Logarithmic in the size of N" = "Proportional to the logarithm of the size of N"

Something about "giving semantic meaning to abstract memory"

"Below, above etc, no... use Fig numbers"

