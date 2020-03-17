import Data.List
import Data.Char

doubleMe x = x + x 
doubleUs x y = x*2 + y*2 
con = "It's a-me, Conan O'Brien!"

reverselist::[Int] -> [Int]
reverselist = \list ->
	case list of
	[] -> []
	x:xs ->reverselist xs ++ [x]

insertsorted::[Int] -> Int -> [Int]
insertsorted = \list -> \v ->
	case list of
		[] -> [v]
		x:xs | v>x -> x:insertsorted xs v
		_->v:list

splitfunc::[Int] -> ([Int],[Int])
splitfunc = \list ->
	case list of
		[] -> ([],[])
		x:xs -> let (e,o) = splitfunc xs
			in if x `mod` 2 == 0 then (x:e,o) else (e,x:o)

splittwo::[Int] ->([Int],[Int])
splittwo = \list ->
	case list of 
		[] ->  ([],[])
		[x] -> ([x],[])
		x:y:zs -> (x:xs,y:ys)
			where (xs,ys) = splittwo zs

merge :: [Int] -> [Int] -> [Int]
merge = \sorted1 -> \sorted2 ->
  case (sorted1,sorted2) of
    ([],[]) -> []
    (x,[]) -> x
    ([],y) -> y
    (x:xs,y:ys) | x<y -> x:merge xs sorted2
    			| x>=y -> y:merge sorted1 ys


-- split into odd/even indices
split :: [Int] -> ([Int], [Int])
split = \list ->
  case list of
    [] -> ([], [])
    x:y:zs -> (x:odd,y:even)
    	where (odd,even) = split zs

-- split, recurse, and merge
mergesort :: [Int] -> [Int]
mergesort = \list ->
  case list of
    [] -> []
    [x] -> [x]
    _ -> merge (mergesort (fst (split list))) (mergesort( snd (split list)))
    
compareWithHundred :: (Num a, Ord a) => a -> Ordering
compareWithHundred = compare 100

compares::(Num a, Ord a) => a -> a-> Bool
compares = \x -> \y ->
	x==y
-----------------------------------------------
-- fold as discussed in class
foldr' :: (a->b->b) -> b -> [a] -> b
foldr' = \fn -> \z -> \list ->
  case list of 
    [] -> z 
    x:xs -> fn x (foldr' fn z xs)

-----------------------------------------
isFunc :: String -> Bool
isFunc = \str ->
  case str of
  "+" -> True
  "-" -> True
  "*" -> True
  _ -> False
  
stringToFunc :: String -> (Int -> Int -> Int)
stringToFunc = \str ->
  case str of
  "-" -> (-)
  "*" -> (*)
  _ -> (+)

foldAll :: [Int] -> String -> [Int]  
foldAll = \list -> \str ->
  case list of
  x:y:rest | isFunc str -> ((stringToFunc str) y x):rest
  _ -> [read str] ++ list

solveRPN :: String -> Int
solveRPN = \str -> 
  head (foldl foldAll [] (words str))

---------------------------------------------------------------------------------------------
--Assignment--

finder::[String] -> [(String, [String])] -> Bool
finder = \input -> \hy ->
	case hy of
		[] -> False
		x:xs -> if ((fst x):[] == input) then True else finder input xs 
		--x:xs | (fst x):[] == input -> True
		--	|otherwise finder input xs

lineHelper::[(String, [String])] -> [String] -> Int -> ([String], [String])
lineHelper = \(y:ys) -> \list -> \len ->
	case list of 
		[] -> ([],[])
		--x:xs | len -> head(snd(splitLine list len))



splitLine :: [String] -> Int -> ([String], [String])
splitLine = \str -> \i ->
	case str of
		[] -> ([],[])
		x:xs -> let(one,two) = splitLine xs (i-length x)
			in if i > length x then (x:one,two) else (one,x:two)
		
enHyp = [("creative", ["cr","ea","ti","ve"]), ("controls", ["co","nt","ro","ls"]), ("achieve", ["ach","ie","ve"]), ("future", ["fu","tu","re"]), ("present", ["pre","se","nt"]), ("motivated", ["mot","iv","at","ed"]), ("desire", ["de","si","re"]), ("others", ["ot","he","rs"])]


spacer:: Int -> [String] -> [String]
spacer = \num -> \list ->
	case num of
		0 -> []
--		 ->  

--blankInsertions :: Int -> [String] -> [[String]]
--blankInsertions = \num -> \list ->
--	case list of 
--		x:xs -> 
func:: [(String, [String])] -> [String] -> [String] 
func = \nq -> \input ->
	case nq of 
		[] -> []
		x:xs | fst x == head input -> snd x
		_:xs -> func xs input  

xtra::[String] -> String
xtra = \list ->
	case list of 
		[]->[]
		x:xs -> x 

lineBreaks :: [(String, [String])] -> Int -> [String] -> [([String], [String])]
lineBreaks = \hy -> \wd -> \list ->
 case list of 
  [] -> []
  x:xs -> [] 


merger::[String] -> [(String,String)]
merger = \list ->
 case list of
  [] -> []
  [x] -> [(x,"")]
  x:y:xs ->[(x++"-",y++(unwords xs))]  ++ merger ((concat [x ,y]):xs)