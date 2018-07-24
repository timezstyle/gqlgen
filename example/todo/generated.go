// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package todo

import (
	"bytes"
	context "context"
	strconv "strconv"

	graphql "github.com/vektah/gqlgen/graphql"
	introspection "github.com/vektah/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser"
	ast "github.com/vektah/gqlparser/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(resolvers ResolverRoot) graphql.ExecutableSchema {
	return &executableSchema{resolvers: resolvers}
}

type ResolverRoot interface {
	MyMutation() MyMutationResolver
	MyQuery() MyQueryResolver
}
type MyMutationResolver interface {
	CreateTodo(ctx context.Context, todo TodoInput) (Todo, error)
	UpdateTodo(ctx context.Context, id int, changes map[string]interface{}) (*Todo, error)
}
type MyQueryResolver interface {
	Todo(ctx context.Context, id int) (*Todo, error)
	LastTodo(ctx context.Context) (*Todo, error)
	Todos(ctx context.Context) ([]Todo, error)
}

type executableSchema struct {
	resolvers ResolverRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._MyQuery(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:   buf,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, op *ast.OperationDefinition) *graphql.Response {
	ec := executionContext{graphql.GetRequestContext(ctx), e.resolvers}

	buf := ec.RequestMiddleware(ctx, func(ctx context.Context) []byte {
		data := ec._MyMutation(ctx, op.SelectionSet)
		var buf bytes.Buffer
		data.MarshalGQL(&buf)
		return buf.Bytes()
	})

	return &graphql.Response{
		Data:   buf,
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, op *ast.OperationDefinition) func() *graphql.Response {
	return graphql.OneShot(graphql.ErrorResponse(ctx, "subscriptions are not supported"))
}

type executionContext struct {
	*graphql.RequestContext

	resolvers ResolverRoot
}

var myMutationImplementors = []string{"MyMutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _MyMutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, myMutationImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "MyMutation",
	})

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MyMutation")
		case "createTodo":
			out.Values[i] = ec._MyMutation_createTodo(ctx, field)
		case "updateTodo":
			out.Values[i] = ec._MyMutation_updateTodo(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _MyMutation_createTodo(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 TodoInput
	if tmp, ok := field.Args["todo"]; ok {
		var err error
		arg0, err = UnmarshalTodoInput(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["todo"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "MyMutation"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.resolvers.MyMutation().CreateTodo(ctx, args["todo"].(TodoInput))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(Todo)
	return ec._Todo(ctx, field.Selections, &res)
}

func (ec *executionContext) _MyMutation_updateTodo(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := field.Args["id"]; ok {
		var err error
		arg0, err = graphql.UnmarshalInt(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["id"] = arg0
	var arg1 map[string]interface{}
	if tmp, ok := field.Args["changes"]; ok {
		var err error
		arg1 = tmp.(map[string]interface{})
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["changes"] = arg1
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "MyMutation"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.resolvers.MyMutation().UpdateTodo(ctx, args["id"].(int), args["changes"].(map[string]interface{}))
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*Todo)
	if res == nil {
		return graphql.Null
	}
	return ec._Todo(ctx, field.Selections, res)
}

var myQueryImplementors = []string{"MyQuery"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _MyQuery(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, myQueryImplementors)

	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "MyQuery",
	})

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("MyQuery")
		case "todo":
			out.Values[i] = ec._MyQuery_todo(ctx, field)
		case "lastTodo":
			out.Values[i] = ec._MyQuery_lastTodo(ctx, field)
		case "todos":
			out.Values[i] = ec._MyQuery_todos(ctx, field)
		case "__type":
			out.Values[i] = ec._MyQuery___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._MyQuery___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _MyQuery_todo(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := field.Args["id"]; ok {
		var err error
		arg0, err = graphql.UnmarshalInt(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["id"] = arg0
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "MyQuery",
		Args:   args,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.MyQuery().Todo(ctx, args["id"].(int))
		})
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.(*Todo)
		if res == nil {
			return graphql.Null
		}
		return ec._Todo(ctx, field.Selections, res)
	})
}

func (ec *executionContext) _MyQuery_lastTodo(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "MyQuery",
		Args:   nil,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.MyQuery().LastTodo(ctx)
		})
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.(*Todo)
		if res == nil {
			return graphql.Null
		}
		return ec._Todo(ctx, field.Selections, res)
	})
}

func (ec *executionContext) _MyQuery_todos(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	ctx = graphql.WithResolverContext(ctx, &graphql.ResolverContext{
		Object: "MyQuery",
		Args:   nil,
		Field:  field,
	})
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.Recover(ctx, r)
				ec.Error(ctx, userErr)
				ret = graphql.Null
			}
		}()

		resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
			return ec.resolvers.MyQuery().Todos(ctx)
		})
		if resTmp == nil {
			return graphql.Null
		}
		res := resTmp.([]Todo)
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler {
				rctx := graphql.GetResolverContext(ctx)
				rctx.PushIndex(idx1)
				defer rctx.Pop()
				return ec._Todo(ctx, field.Selections, &res[idx1])
			}())
		}
		return arr1
	})
}

func (ec *executionContext) _MyQuery___type(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := field.Args["name"]; ok {
		var err error
		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["name"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "MyQuery"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.introspectType(args["name"].(string)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) _MyQuery___schema(ctx context.Context, field graphql.CollectedField) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "MyQuery"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return ec.introspectSchema(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	if res == nil {
		return graphql.Null
	}
	return ec.___Schema(ctx, field.Selections, res)
}

var todoImplementors = []string{"Todo"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Todo(ctx context.Context, sel ast.SelectionSet, obj *Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, todoImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			out.Values[i] = ec._Todo_id(ctx, field, obj)
		case "text":
			out.Values[i] = ec._Todo_text(ctx, field, obj)
		case "done":
			out.Values[i] = ec._Todo_done(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Todo_id(ctx context.Context, field graphql.CollectedField, obj *Todo) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Todo"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.ID, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(int)
	return graphql.MarshalInt(res)
}

func (ec *executionContext) _Todo_text(ctx context.Context, field graphql.CollectedField, obj *Todo) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Todo"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Text, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) _Todo_done(ctx context.Context, field graphql.CollectedField, obj *Todo) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "Todo"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Done, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(bool)
	return graphql.MarshalBoolean(res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __DirectiveImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Locations, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]string)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return graphql.MarshalString(res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Directive"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Args, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __EnumValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(bool)
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__EnumValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __FieldImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Args, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Type, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.IsDeprecated, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(bool)
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Field"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.DeprecationReason, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __InputValueImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Type, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__InputValue"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.DefaultValue, nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __SchemaImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Types(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.QueryType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.MutationType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.SubscriptionType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Schema"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Directives(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Directive)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Directive(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ctx, sel, __TypeImplementors)

	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Kind(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Name(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Description(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Fields(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Field)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Field(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.Interfaces(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.PossibleTypes(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.Type)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___Type(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error
		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(ctx, err)
			return graphql.Null
		}
	}
	args["includeDeprecated"] = arg0
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = args
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.EnumValues(args["includeDeprecated"].(bool)), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.EnumValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___EnumValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.InputFields(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]introspection.InputValue)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			rctx := graphql.GetResolverContext(ctx)
			rctx.PushIndex(idx1)
			defer rctx.Pop()
			return ec.___InputValue(ctx, field.Selections, &res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	rctx := graphql.GetResolverContext(ctx)
	rctx.Object = "__Type"
	rctx.Args = nil
	rctx.Field = field
	rctx.PushField(field.Alias)
	defer rctx.Pop()
	resTmp := ec.FieldMiddleware(ctx, func(ctx context.Context) (interface{}, error) {
		return obj.OfType(), nil
	})
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func UnmarshalTodoInput(v interface{}) (TodoInput, error) {
	var it TodoInput
	var asMap = v.(map[string]interface{})

	for k, v := range asMap {
		switch k {
		case "text":
			var err error
			it.Text, err = graphql.UnmarshalString(v)
			if err != nil {
				return it, err
			}
		case "done":
			var err error
			var ptr1 bool
			if v != nil {
				ptr1, err = graphql.UnmarshalBoolean(v)
				it.Done = &ptr1
			}

			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name])
}

var parsedSchema = gqlparser.MustLoadSchema(
	&ast.Source{Name: "schema.graphql", Input: `schema {
	query: MyQuery
	mutation: MyMutation
}

type MyQuery {
	todo(id: Int!): Todo
	lastTodo: Todo
	todos: [Todo!]!
}

type MyMutation {
	createTodo(todo: TodoInput!): Todo!
	updateTodo(id: Int!, changes: Map!): Todo
}

type Todo {
	id: Int!
	text: String!
	done: Boolean!
}

input TodoInput {
	text: String!
	done: Boolean
}

scalar Map
`},
)
