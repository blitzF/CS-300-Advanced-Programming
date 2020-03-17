import Data.List
import Data.Char

enHyp = [("creative", ["cr","ea","ti","ve"]), ("controls", ["co","nt","ro","ls"]), ("achieve", ["ach","ie","ve"]), ("future", ["fu","tu","re"]), ("present", ["pre","se","nt"]), ("motivated", ["mot","iv","at","ed"]), ("desire", ["de","si","re"]), ("others", ["ot","he","rs"])]

merger::[String] -> [(String,String)]
merger = \list  ->
 case list of
  [] -> []
  [x]  -> [(x,"")]
  x:y:xs ->[(x++"-",y++(concat xs))]  ++ merger ((concat [x ,y]):xs)

lineBreaks :: [(String, [String])] -> Int -> [String] -> [([String], [String])]
lineBreaks = \hy -> \wd -> \list ->
 case list of 
  [] -> []
  x:xs -> let (one,two) = splitLine list wd
   in if (finder (head two) enHyp == True) then filter(\(x,_) -> ((length x) + sum(map(\y -> length y) x) -1 ) <= wd)  ((one,two):map(\(x,y)->((one++[x]),[y]++(tail two)))(merger(func hy two))) else [(one,two)]

splitLine :: [String] -> Int -> ([String], [String])
-- Function definition here
splitLine = \str -> \i ->
 case str of
  [] -> ([],[])
  x:xs -> let(one,two) = splitLine xs (i-length x)
   in if i > length x then (x:one,two) else (one,x:two)


func:: [(String, [String])] -> [String] -> [String] 
func = \nq -> \input ->
	case nq of 
		[] -> []
		x:xs | fst x == head input -> snd x
		_:xs -> func xs input

finder::String -> [(String, [String])] -> Bool
finder = \input -> \hy ->
	case hy of
		[] -> False
		x:xs -> if ((fst x) == input) then True else finder input xs

blanker :: [String] -> [[String]]
blanker = \list ->
 case list of
 	[] -> [[]] 
 	[x,y] -> [[x," ",y]]
 	x:xs -> [x:" ":xs] ++ (map (\rest -> [x]++rest ) (blanker xs))

recurblanker::Int -> [[String]] ->[[String]]
recurblanker = \val -> \list ->
 case val of
 	0 -> list
 	_-> recurblanker (val-1) (concat(map(blanker)list)) 

mur::[String] -> [[String]] ->[[String]]
mur = \li -> \lii -> [concat(li:lii)]

uniq :: [[String]] -> [[String]] -> [[String]]
uniq = \emp -> \input ->
 case input of
 	[]->emp
 	x:xs -> if (listcheck_list emp x) == True then uniq emp xs else uniq(x:emp) xs 

listcheck::[String] -> [String] -> Bool
listcheck = \li -> \lii ->
 case li of
 	[] | (length lii) == 0  -> True
 	x:xs | (length lii) == 0  -> False
 	x:xs | x==(head lii) -> True && (listcheck xs (tail lii))
 	_-> False


listcheck_list::[[String]] -> [String] -> Bool
listcheck_list = \li -> \lii ->
 case li of
 	[] -> False
 	x:xs | (listcheck x lii)==True -> True || (listcheck_list xs lii)
 	_ -> False || (listcheck_list (tail li) lii)

blankInsertions :: Int -> [String] -> [[String]]
blankInsertions = \val -> \list -> (uniq [] (recurblanker val [list]))

 