package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"yooplus_indication/models"
)

type ArticleController struct {
	BaseController
}

//测试页面
func (this *ArticleController) Get() {
	this.TplName = "article_test.html"
}

//////////////////////////////////////////////////////////////////////
//																	//
//							亲友圈接口				                //
//																	//
//////////////////////////////////////////////////////////////////////

// options == 0 查询亲友圈列表
// options == 1 查询指定的亲友圈信息
// options == 2 发表亲友圈消息
// options == 3 删除指定亲友圈消息
// options == 4 发表评论
// options == 5 删除评论

func (this *ArticleController) Post() {

	userId, _ := strconv.ParseInt(this.Input().Get("userId"), 10, 64)
	options, _ := strconv.ParseInt(this.Input().Get("options"), 10, 64)
	beego.Info(userId)
	beego.Info(options)

	// options == 0 查询亲友圈列表
	if options == 0 {
		articles, err := this.getArticles(userId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getArticles error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		comments, err := this.getCommentsByUserId(userId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getComments error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "articles": articles, "comments": comments, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	// options == 1 查询指定的亲友圈信息
	if options == 1 {
		articleId, _ := strconv.ParseInt(this.Input().Get("articleId"), 10, 64)
		beego.Info(articleId)
		article, err := this.getArticle(articleId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getArticle error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		comments, err := this.getCommentsByArticleId(articleId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "getComments error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "article": article, "comments": comments, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	// options == 2 发表亲友圈消息
	if options == 2 {
		desc := this.Input().Get("desc")
		createTime := this.Input().Get("createTime")
		articleId, err := this.addArticle(userId, desc, createTime)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "addArticle error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "articleId": articleId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	// options == 3 删除指定亲友圈消息
	if options == 3 {
		articleId, _ := strconv.ParseInt(this.Input().Get("articleId"), 10, 64)
		err := this.delArticle(articleId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delArticle error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		err = this.delCommentByArticleId(articleId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delComments error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		err = this.delPhotosByArticleId(articleId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delPhotos error", "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "delArticle success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	// options == 4 发表评论
	if options == 4 {
		articleId, _ := strconv.ParseInt(this.Input().Get("articleId"), 10, 64)
		commentId, _ := strconv.ParseInt(this.Input().Get("commentId"), 10, 64)
		desc := this.Input().Get("commentId")
		// commentId == -1,针对article的评论
		mCommentId, err := this.addComment(userId, articleId, commentId, desc)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "addComment error --- " + err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "commentId": mCommentId, "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return
	}

	// options == 5 删除评论
	if options == 5 {
		commentId, _ := strconv.ParseInt(this.Input().Get("commentId"), 10, 64)
		err := this.delComments(commentId)
		if err != nil {
			beego.Error(err)
			this.Data["json"] = map[string]interface{}{"status": 400, "message": "delComments error --- " + err.Error(), "time": time.Now().Format("2006-01-02 15:04:05")}
			this.ServeJSON()
			return
		}
		this.Data["json"] = map[string]interface{}{"status": 200, "message": "del comment success", "time": time.Now().Format("2006-01-02 15:04:05")}
		this.ServeJSON()
		return

	}

	this.Data["json"] = map[string]interface{}{"status": 400, "message": "options error", "time": time.Now().Format("2006-01-02 15:04:05")}
	this.ServeJSON()
	return

}

//////////////////////////////////////////////////////////////////////
//																	//
//				将数据库操作再封装一层方便替换model层               //
//																	//
//////////////////////////////////////////////////////////////////////

func (this *ArticleController) getArticles(userId int64) ([]*models.TArticle, error) {
	articles, err := models.GetArticlesByUserId(userId)
	return articles, err
}

func (this *ArticleController) getArticle(articleId int64) (*models.TArticle, error) {
	article, err := models.GetArticleById(articleId)
	return article, err
}

func (this *ArticleController) addArticle(userId int64, desc, createTime string) (int64, error) {
	article := &models.TArticle{Dcsc: desc, UserId: userId, CreateTime: createTime}
	articleId, err := models.AddArticle(article)
	return articleId, err
}

func (this *ArticleController) delArticle(articleId int64) error {
	article := &models.TArticle{Id: articleId}
	err := models.DelArticle(article)
	return err
}

func (this *ArticleController) getCommentsByUserId(userId int64) ([]*models.TComment, error) {
	comments, err := models.GetCommentsByUserId(userId)
	return comments, err
}

func (this *ArticleController) getCommentsByArticleId(articleId int64) ([]*models.TComment, error) {
	comments, err := models.GetCommentsByArticleId(articleId)
	return comments, err
}

func (this *ArticleController) addComment(userId, articleId, commentId int64, desc string) (int64, error) {
	comment := &models.TComment{UserId: userId, ArticleId: articleId, CommentId: commentId, Desc: desc, CreateTime: time.Now().Format("2006-01-02 15:04:05")}
	mCommentId, err := models.AddComment(comment)
	return mCommentId, err
}

func (this *ArticleController) delComments(commentId int64) error {
	comment := &models.TComment{Id: commentId}
	err := models.DelComment(comment)
	return err
}

func (this *ArticleController) delCommentByArticleId(articleId int64) error {
	comments, err := models.GetCommentsByArticleId(articleId)
	if err != nil {
		return err
	}
	for i := 0; i < len(comments); i++ {
		err := models.DelComment(comments[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *ArticleController) delPhotosByArticleId(articleId int64) error {
	photos, err := models.GetPhotosByArticleId(articleId)
	if err != nil {
		return err
	}
	for i := 0; i < len(photos); i++ {
		err := models.DelPhoto(photos[i])
		if err != nil {
			return err
		}
	}
	return nil
}
