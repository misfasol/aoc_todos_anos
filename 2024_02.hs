import Data.List
import Data.Ord
import System.IO

-- funcao1 :: String -> Int
funcao1 x =
  let linhas = map (map (\b -> read b :: Int) . words) (lines x)
      sorted = filter (\a -> a == sort a || a == sortBy (comparing Data.Ord.Down) a) linhas
      entre = filter (\a -> maiMenDif a max <= 3 && maiMenDif a min >= 1) sorted
   in length entre

maiMenDif :: [Int] -> (Int -> Int -> Int) -> Int
maiMenDif [] _ = 0
maiMenDif [x] _ = 0
maiMenDif [x, y] _ = abs (x - y)
maiMenDif (x : y : xs) op = op (abs (x - y)) (maiMenDif (y : xs) op)

funcao2 :: String -> Int
funcao2 x =
  let linhas = map (map (\b -> read b :: Int) . words) (lines x)
      sorted = filter (\a -> a == sort a || a == sortBy (comparing Data.Ord.Down) a) linhas
      entre = filter (\a -> maiMenDif a max <= 3 && maiMenDif a min >= 1) sorted
   in length entre

main = do
  arq <- openFile "2024_02.txt" ReadMode
  conteudo <- hGetContents arq
  print $ funcao1 conteudo
  print $ funcao2 conteudo
