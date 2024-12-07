import Control.Monad
import System.IO

funcao1 :: [[String]] -> Int
funcao1 l = sum $ map ((\a -> read a :: Int) . init . head) (filter valido1 l)

valido1 :: [String] -> Bool
valido1 (x : xs) =
  let valor = read (init x) :: Int
      numes = map (\a -> read a :: Int) xs
      opera = replicateM (length numes - 1) [(*), (+)]
      resul = map (avaliarOpe1 numes) opera
   in elem valor resul

avaliarOpe1 :: [Int] -> [Int -> Int -> Int] -> Int
avaliarOpe1 [n1, n2] [o] = o n1 n2
avaliarOpe1 (n : ns : nx) (o : os) = avaliarOpe1 (o n ns : nx) os

funcao2 :: [[String]] -> Int
funcao2 l = sum $ map ((\a -> read a :: Int) . init . head) (filter valido2 l)

valido2 :: [String] -> Bool
valido2 (x : xs) =
  let valor = read (init x) :: Int
      numes = map (\a -> read a :: Int) xs
      opera = replicateM (length numes - 1) [(*), (+), conc]
      resul = map (avaliarOpe2 numes) opera
   in elem valor resul

conc :: Int -> Int -> Int
conc a b = read (show a ++ show b) :: Int

avaliarOpe2 :: [Int] -> [Int -> Int -> Int] -> Int
avaliarOpe2 [n1, n2] [o] = o n1 n2
avaliarOpe2 (n : ns : nx) (o : os) = avaliarOpe2 (o n ns : nx) os

main :: IO ()
main = do
  arq <- openFile "2024_07.txt" ReadMode
  conteudo <- hGetContents arq
  let linhas = map words $ lines conteudo
  print $ funcao1 linhas
  print $ funcao2 linhas
