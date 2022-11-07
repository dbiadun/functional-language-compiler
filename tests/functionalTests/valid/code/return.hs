main :: IO ();
main = do {
    x <- return523;
    putInt x;
    putStr "\n";
};

return523 :: IO Int;
return523 = return 523;