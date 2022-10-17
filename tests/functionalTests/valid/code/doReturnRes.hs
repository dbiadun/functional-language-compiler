main :: IO ();
main = do {
    x <- returnRes;
    putStr x;
};

returnRes :: IO String;
returnRes = do {
    putStr "returnRes\n";
    return "something";
};
