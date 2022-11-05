import "List.hs";
import "../../libs/SortType.hs";

data Tree a = NoNode | Node Int a (Tree a) (Tree a); -- The integer value tells what is the distance to the closest leaf
data HeapAndList a = HeapAndList (Tree a) (List a);

main :: IO ();
main = do {
    tinySleep 2000;
    
    l <- return (getList 16);
    l <- return (heapSort l);

    printList l;
    putStr "\n";
};

heapSort :: List SortType -> List SortType;
heapSort l = getListFromHL (heapToList (HeapAndList (buildHeap l) Nil));

getListFromHL :: HeapAndList a -> List a;
getListFromHL hl = case hl of {
    HeapAndList h l -> l;
};

heapToList :: HeapAndList SortType -> HeapAndList SortType;
heapToList hl = case hl of {
    HeapAndList h list -> case h of {
        NoNode -> HeapAndList h list;
        Node d v l r -> heapToList (HeapAndList (removeOne h) (Cons v list));
    };
};

removeOne :: Tree SortType -> Tree SortType;
removeOne t = case t of {
    NoNode -> t;
    Node d v l r -> case l of {
        NoNode -> case r of {
            NoNode -> NoNode;
            x -> x;
        };
        Node dl vl ll rl -> case r of {
            NoNode -> l;
            Node dr vr lr rr -> case cmp vr vl of {
                True -> linkTrees vl (removeOne l) r;
                False -> linkTrees vr l (removeOne r);
            };
        };
    };
};

not :: Bool -> Bool;
not b = case b of {
    True -> False;
    False -> True;
};

buildHeap :: List SortType -> Tree SortType;
buildHeap = foldl addNode NoNode;

addNode :: Tree SortType -> SortType -> Tree SortType;
addNode h el = case h of {
    NoNode -> Node 0 el NoNode NoNode;
    Node d v l r -> sortNodeAndChildren (addNodeToSubtree el v l r);
};

addNodeToSubtree :: SortType -> SortType -> Tree SortType -> Tree SortType -> Tree SortType;
addNodeToSubtree el v l r = case getDepth l <= getDepth r of {
    True -> linkTrees v (addNode l el) r;
    False -> linkTrees v l (addNode r el);
};

linkTrees :: SortType -> Tree SortType -> Tree SortType -> Tree SortType;
linkTrees v l r = Node ((getDepth r) + 1) v l r;

sortNodeAndChildren :: Tree SortType -> Tree SortType;
sortNodeAndChildren t = case t of {
    NoNode -> NoNode;
    Node d v l r -> case shouldSwap v l of {
        True -> Node d (getVal l) (setVal v l) r;
        False -> case shouldSwap v r of {
            True -> Node d (getVal r) l (setVal v r);
            False -> t;
        };
    };
};

shouldSwap :: SortType -> Tree SortType -> Bool;
shouldSwap v t = case t of {
    NoNode -> False;
    Node d v1 l r -> cmp v v1;
};

getDepth :: Tree a -> Int;
getDepth t = case t of {
    NoNode -> -1;
    Node d v left right -> d;
};

getVal :: Tree a -> a;
getVal t = case t of {
    Node d v left right -> v;
};

setVal :: a -> Tree a -> Tree a;
setVal v t = case t of {
    Node d v0 l r -> Node d v l r;
};
