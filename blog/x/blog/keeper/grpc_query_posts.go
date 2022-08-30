package keeper

import (
  "context"
  
  "github.com/cosmos/cosmos-sdk/store/prefix"
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/cosmos/cosmos-sdk/types/query"  
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"

  "blog/x/blog/types"
)

func (k Keeper) Posts(c context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	// Выдать ошибку, если запрос равен нулю
  if req == nil {
    return nil, status.Error(codes.InvalidArgument, "invalid request")
  }

  // Определяем переменную, в которой будет храниться список постов
  var posts []*types.Post

  // Получаем контекст с информацией об окружающей среде
  ctx := sdk.UnwrapSDKContext(c)

  // Получите хранилище модуля "ключ-значение" с помощью ключа хранилища (в нашем случае ключом хранилища является "цепочка").
  store := ctx.KVStore(k.storeKey)

  // Получить часть хранилища, которая хранит сообщения (используя ключ сообщения, который равен «Post-value-»)
  postStore := prefix.NewStore(store, []byte(types.PostKey))

  // Разбивка хранилища сообщений на страницы на основе PageRequest
  pageRes, err := query.Paginate(postStore, req.Pagination, func(key []byte, value []byte) error {
    var post types.Post
    if err := k.cdc.Unmarshal(value, &post); err != nil {
      return err
    }

    posts = append(posts, &post)

    return nil
  })

  // Throw an error if pagination failed
  if err != nil {
    return nil, status.Error(codes.Internal, err.Error())
  }

  // Возвращаем структуру, содержащую список постов и информацию о пагинации
  return &types.QueryPostsResponse{Post: posts, Pagination: pageRes}, nil
}