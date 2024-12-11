import Control.Monad
import System.IO

funcao1 l = funcao1H (map (\a -> read a :: Integer) l) 0 25

funcao1H l qtd vzs
  | qtd == vzs = length l
  | otherwise = funcao1H (l >>= regra1) (qtd + 1) vzs

regra1 :: Integer -> [Integer]
regra1 n
  | n == 0 = [1]
  | even (qtdDig n) = [n `div` (10 ^ (qtdDig n `div` 2)), n `mod` (10 ^ (qtdDig n `div` 2))]
  | otherwise = [n * 2024]

qtdDig :: Integer -> Int
qtdDig n
  | n == 0 = 0
  | n < 10 = 1
  | otherwise = 1 + qtdDig (n `div` 10)

funcao2 l = funcao1H (map (\a -> read a :: Integer) l) 0 75

main :: IO ()
main = do
  arq <- openFile "2024_11.txt" ReadMode
  conteudo <- hGetContents arq
  let linhas = words conteudo
  print $ funcao1 linhas
  print $ funcao2 linhas