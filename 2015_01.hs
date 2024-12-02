import System.IO

funcao :: String -> Int
funcao x = length (filter (== '(') x) - length (filter (== ')') x)

main :: IO ()
main = do
  arq <- openFile "2015_01.txt" ReadMode
  conteudo <- hGetContents arq
  print $ funcao conteudo
