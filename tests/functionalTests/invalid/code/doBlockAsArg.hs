main :: IO ();
main = do {
    x <- return 14;
    putInt x;
    putStr "\n";
    runAction (do {
       putStr "do block\n";
       x <- return 10;
       putInt x;
       putStr "\n";
    });
    putInt x;
    putStr "\n";
};

runAction :: IO () -> IO ();
runAction action = do {
    putStr "runAction\n";
    action;
    putStr "end runAction\n";
};
