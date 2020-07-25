package dao

import "github.com/cnodin/blog-service/internal/model"

func (d *Dao) GetArticleTagByAID(articleID uint32) (model.ArticleTag, error) {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	return articleTag.GetByID(d.engine)
}

func (d *Dao) GetArticleTagListByTID(tagID uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{
		TagID: tagID,
	}
	return articleTag.ListByTID(d.engine)
}

func (d *Dao) GetArticleTagListByAIDs(articleIDs []uint32) ([]*model.ArticleTag, error) {
	articleTag := model.ArticleTag{}
	return articleTag.ListByAIDs(d.engine, articleIDs)
}

func (d *Dao) CreateArticleTag(articleID, tagID uint32, createBy string) error {
	articleTag := model.ArticleTag{
		Model:     &model.Model{
			CreatedBy: createBy,
		},
		TagID:     tagID,
		ArticleID: articleID,
	}
	return articleTag.Create(d.engine)
}

func (d *Dao) UpdateArticleTag(articleID, tagID uint32, modifiedBy string) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	values := map[string]interface{}{
		"article_id": articleID,
		"tag_id": tagID,
		"modified_by": modifiedBy,
	}
	return articleTag.Update(d.engine, values)
}

func (d *Dao) DeleteArticleTag(articleID uint32) error {
	articleTag := model.ArticleTag{
		ArticleID: articleID,
	}
	return articleTag.DeleteOne(d.engine)
}
