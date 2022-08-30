package keeper

import (
	"blog/x/blog/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// получаем текущее количество постов в хранилище
	count := k.GetPostCount(ctx)
	// Присвоить ID сообщению на основе количества сообщений в магазине
	post.Id = count
	// получаем обьект хранилища
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	// создаём массив байтов и записываем туда количество сообщений в хранилище
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)
	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&post)
	//// Вставляем байты поста, используя идентификатор поста в качестве ключа
	store.Set(byteKey, appendedValue)

	//обновляем счётчик поста
	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	// Преобразование PostCountKey в байты PostCountKey (это «Post-count-»)
	byteKey := []byte(types.PostCountKey)

	// Получаем обьект хранилища используюя storeKey (это «блог») и PostCountKey (это «Post-count-»)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	// с помощью обьекта хранилища получаем и его метода Get получаем счётчик
	bz := store.Get(byteKey)
	// Возвращаем ноль, если значение счетчика не найдено (например, это первый пост)
	if bz == nil {
		return 0
	}
	// конвертируем счётчик в юинт64
	return binary.BigEndian.Uint64(bz)

}

func (k Keeper) SetPostCount(ctx sdk.Context, counter uint64) {
	byteKey := []byte(types.PostCountKey)

	// получаем обьект хранилища с помощью ключей
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	//Конвертируем count из uint64 в строку и получаем байты
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, counter)

	// Установить значение Post-count- для подсчета
	store.Set(byteKey, bz)
}
