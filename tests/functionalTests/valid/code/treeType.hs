data Tree a = Node a (Tree a) (Tree a) | NoNode;

printIntTree :: Tree Int -> IO ();
printIntTree t = case t of {
    Node num left right -> do {
        putInt num;

        putStr " (";
        printIntTree left;
        putStr ")";

        putStr " (";
        printIntTree right;
        putStr ")";
    };
    NoNode -> putStr "";
};

main :: IO ();
main = do {
    tree <- return (Node 2 (Node 5 (Node 7 (NoNode) (Node 15 (NoNode) (NoNode))) (NoNode)) (Node 20 (Node 15 (NoNode) (NoNode)) (Node 16 (NoNode) (NoNode))));
    printIntTree tree;
};