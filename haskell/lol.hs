numblanks::[String] -> Double
numblanks = \list ->
 case list of
  []->0
  x:xs | x == " " -> 1.0 + numblanks xs
  _->numblanks (tail list)

blankdist::[String] -> Int
blankdist = \list -> 
 case list of
  []->0
  x:xs | x==" " -> 0
  _-> 1+ blankdist (tail list)

blankdistlist::[String] -> Int->[Int]
blankdistlist = \list -> \val -> 
 case list of
 []->val:[]
 x:xs | x==" "-> val:[] ++ (blankdistlist xs 0)
 _->blankdistlist (tail list) (val+1)

numlist::[String]->[Int]
numlist = \list ->(blankdist list):[]

blankchecker::[String]->Bool
blankchecker = \list ->
 case list of
  []->False
  x:xs | x==" " -> True
  _->False

avg::[Int]->Double
avg = \list -> (fromIntegral(sum list)) / (fromIntegral (length list))
--sum xs / (fromIntegral (length xs))

data Tree a = EmptyTree | Node a (Tree a) (Tree a) deriving (Show,Read,Eq)

treeins::(Ord a) => a -> Tree a -> Tree a
treeins = \val -> \tr ->
 case tr of
  EmptyTree -> Node val EmptyTree EmptyTree
  (Node x left right) | val == x -> Node x left right
  (Node x left right) | val < x -> Node x (treeins val left) right
  (Node x left right) | val > x -> Node x left (treeins val right)

mytree :: Tree Int
mytree = EmptyTree


iom::Int->Int->Char->Char
iom = \r -> \c -> \ch ->
 case (r,c) of
  (0,0) -> ch