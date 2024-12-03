package node

// // A Registry is a registry of nodes which stores node
// // data by inferring type associations to the node's outputs
// // and allows for quick lookups of nodes by their outputs.
// type Registry struct {
// 	nodes map[reflect.Type]*Node
// }

// // NewRegistry creates a new NodeRegistry.
// func NewRegistry() *Registry {
// 	return &Registry{
// 		nodes: make(map[reflect.Type]*Node),
// 	}
// }

// // Register adds a node to the registry.
// func (r *Registry) Register(node *Node) error {
// 	for _, output := range node.Outputs() {
// 		if _, ok := r.nodes[output]; ok {
// 			return errors.New(ErrDuplicateOutput, registryErrorName, output.String())
// 		}
// 		r.nodes[output] = node
// 	}

// 	return nil
// }

// // Get retrieves a node from the registry by its output type.
// // TODO: implement list inferencing.
// func (r *Registry) Get(t reflect.Type) (*Node, error) {
// 	// Check if the direct type exists in the registry.
// 	node, ok := r.nodes[t]
// 	if ok {
// 		return node, nil
// 	}

// 	// Check if the type is an interface, and if it is,
// 	// check if there is an implementation registered.
// 	if t.Kind() == stdreflect.Interface {
// 		var foundNode *Node
// 		for regType, node := range r.nodes {
// 			if regType.Implements(t) {
// 				if foundNode != nil {
// 					return nil, ErrMultipleImplementations(t)
// 				}
// 				foundNode = node
// 			}
// 		}
// 		if foundNode != nil {
// 			return foundNode, nil
// 		}
// 	}

// 	return nil, ErrMissingDependency(t)
// }

// // Nodes returns all nodes in the registry.
// func (r *Registry) Nodes() []*Node {
// 	nodes := make([]*Node, 0, len(r.nodes))
// 	for _, node := range r.nodes {
// 		nodes = append(nodes, node)
// 	}
// 	return nodes
// }
