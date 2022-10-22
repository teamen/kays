package mysql

import (
	"context"
	"fmt"
	"log"
	"testing"

	v1 "github.com/teamen/kays/internal/pkg/model/apiserver/v1"
)

func TestCategoryCreateRootNode(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown(t)

	slug := "lens-series"
	rootNode := &v1.Category{
		Title: "镜片系列",
		Slug:  slug,
	}

	log.Printf("%+v", rootNode)

	ctx := context.Background()
	categoryStore := newCategories(ds)
	if err := categoryStore.Create(ctx, rootNode, nil); err != nil {
		t.Fatalf("failed to create root node category: %s", err.Error())
	}

	node, err := categoryStore.GetBySlug(ctx, slug)
	if err != nil {
		t.Fatalf("failed to retrive node by slug[%s]:%v", slug, err)
	}
	t.Logf("%v", node)

}

func TestCategoryCreateChildNode(t *testing.T) {

	tearDown := setUp(t)
	defer tearDown(t)

	slug := "lens-series"
	rootNode := &v1.Category{
		Title: "镜片系列",
		Slug:  slug,
	}

	log.Printf("%+v", rootNode)

	ctx := context.Background()
	categoryStore := newCategories(ds)

	node, err := categoryStore.GetBySlug(ctx, slug)
	if err != nil {
		t.Logf("failed to retrive node by slug[%s]:%v", slug, err)
	}
	t.Logf("%v", node)

	child := &v1.Category{
		Title: "蔡司防蓝光系列",
		Slug:  "lens-series__zeiss-FLG",
	}

	anotherChild := &v1.Category{
		Title: "蔡司A系列",
		Slug:  "lens-series__zeiss-RS",
	}

	if err := categoryStore.Create(ctx, child, node); err != nil {
		t.Logf("failed to create child node :%v", err)
	}

	if err := categoryStore.Create(ctx, anotherChild, node); err != nil {
		t.Logf("failed to create child node :%v", err)
	}

	slug = "product-type"
	productType := &v1.Category{
		Title: "产品类别",
		Slug:  "product-type",
	}

	if err := categoryStore.Create(ctx, productType, nil); err != nil {
		t.Logf("failed to create root node :%v", err)
	}

	fmt.Printf("%+v", productType)

	frame := &v1.Category{
		Title: "镜框",
		Slug:  "product-type__frame",
	}

	lens := &v1.Category{
		Title: "镜片",
		Slug:  "product-type__lens",
	}
	if err := categoryStore.Create(ctx, frame, productType); err != nil {
		t.Logf("failed to create root node :%v", err)
	}

	if err := categoryStore.Create(ctx, lens, productType); err != nil {
		t.Logf("failed to create root node :%v", err)
	}
}

func TestCategoryGetDescendants(t *testing.T) {

	tearDown := setUp(t)
	defer tearDown(t)

	slug := "lens-series"

	ctx := context.Background()
	categoryStore := newCategories(ds)
	node, err := categoryStore.GetBySlug(ctx, slug)
	if err != nil {
		t.Fatalf("failed to retrive the root node which slug is %s:%v", slug, err)
	}

	t.Logf("%+v", node)
	descendants, _ := categoryStore.GetDescendants(ctx, node)
	t.Logf("%+v", descendants)
}

func TestCategoryList(t *testing.T) {

	tearDown := setUp(t)
	defer tearDown(t)

	categoryStore := newCategories(ds)
	ctx := context.Background()

	categoryList, err := categoryStore.List(ctx)
	if err != nil {
		t.Fatalf("failed to list nodes: %v", err)
	}

	t.Logf("%+v", categoryList)

	log.Printf("\n\n")

	for i, c := range categoryList {
		log.Printf("[%d] %+v\n", i, c)
	}
}

func init() {
	log.Println("init...")
}
