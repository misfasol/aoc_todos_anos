import Data.List
import Data.Ord
import System.IO

-- funcao1 :: String -> Int
funcao1 x = length $ filter checar x

checar :: [Int] -> Bool
checar l = (l == sort l || l == sortBy (comparing Data.Ord.Down) l) && maiMenDif l max <= 3 && maiMenDif l min >= 1

maiMenDif :: [Int] -> (Int -> Int -> Int) -> Int
maiMenDif [] _ = 0
maiMenDif [x] _ = 0
maiMenDif [x, y] _ = abs (x - y)
maiMenDif (x : y : xs) op = op (abs (x - y)) (maiMenDif (y : xs) op)

-- funcao2 :: String -> Int
funcao2 mat =
  let filtado = filter removeIndCheca mat
   in length filtado

removeIndCheca :: [Int] -> Bool
removeIndCheca lin = any (checar . removeInd lin) [0 .. length lin - 1]

slice :: Int -> Int -> [a] -> [a]
slice from to xs = take (to - from) (drop from xs)

removeInd :: [Int] -> Int -> [Int]
removeInd l i = slice 0 i l ++ slice (i + 1) (length l) l

main = do
  arq <- openFile "2024_02.txt" ReadMode
  conteudo <- hGetContents arq
  let linhas = map (map (\b -> read b :: Int) . words) (lines conteudo)
  print $ funcao1 linhas
  print $ funcao2 linhas
