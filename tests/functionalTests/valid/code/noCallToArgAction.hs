main :: IO ();
main = do {
    putStr "Action 1\n";
    doNothing1 (putStr "Sample action 1\n");
    putStr "Action 2\n";
    doNothing2 sampleAction;
    putStr "The end\n";
};

doNothing1 :: IO () -> IO ();
doNothing1 x = putStr "";

doNothing2 :: (Int -> IO ()) -> IO ();
doNothing2 x = putStr "";

sampleAction :: Int -> IO ();
sampleAction x = putStr "Sample action 2\n"
