import Data.List
import System.IO

funcao1 :: String -> Int
funcao1 x =
  let linhas = lines x
      duplas = map words linhas
      numeros = map (\[a, b] -> [read a :: Int, read b :: Int]) duplas
      mudado = [map (\[a, b] -> a) numeros, map (\[a, b] -> b) numeros]
      sorted = map sort mudado
      zipad = zip (head sorted) (last sorted)
      difer = map (uncurry (-)) zipad
      abis = map abs difer
   in sum abis

funcao2 :: String -> Int
funcao2 x =
  let linhas = lines x
      duplas = map words linhas
      numeros = map (\[a, b] -> [read a :: Int, read b :: Int]) duplas
      mudado = [map (\[a, b] -> a) numeros, map (\[a, b] -> b) numeros]
      asd1 = map (\a -> length $ filter (== a) (last mudado)) (head mudado)
      asd2 = zip asd1 (head mudado)
      asd3 = sum (map (uncurry (*)) asd2)
   in asd3

main = do
  arq <- openFile "2024_01.txt" ReadMode
  conteudo <- hGetContents arq
  print $ funcao1 conteudo
  print $ funcao2 conteudo
