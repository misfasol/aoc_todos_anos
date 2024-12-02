import System.IO
import Data.Char

funcao :: String -> Int
funcao x = foldr1 (+) (map (\x -> read x :: Int) (map (\x -> [head x, last x]) (map (filter isDigit) (lines x))))

main :: IO ()
main = do
  arq <- openFile "2023_01.txt" ReadMode
  conteudo <- hGetContents arq
  print $ funcao conteudo