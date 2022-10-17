data FibState = FibState Int Int Int;

main :: IO ();
main = do {
    while fibCond fibAction (FibState 0 1 5000000)
};

-- a stores the current state
while :: (a -> IO Bool) -> (a -> IO a) -> a -> IO ();
while cond action state = do {
    condSatisfied <- cond state;
    case condSatisfied of {
        False -> putStr "";
        True -> do {
            newState <- action state;
            while cond action newState;
        };
    };
};

fibCond :: FibState -> IO Bool;
fibCond state = case state of {
    FibState low high max -> return (high < max);
};

fibAction :: FibState -> IO FibState;
fibAction state = case state of {
    FibState low high max -> do {
        putInt high;
        putStr "\n";
        newHigh <- return (low + high);
        newLow <- return high;
        return (FibState newLow newHigh max);
    };
};