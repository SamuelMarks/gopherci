	"context"
func (a *mockAnalyser) NewExecuter(_ context.Context, _ string) (Executer, error) {
func (a *mockAnalyser) Execute(_ context.Context, args []string) (out []byte, err error) {
func (a *mockAnalyser) Stop(_ context.Context) error {
		{ID: 1, Name: "Name1", Path: "tool1", Args: "-flag %BASE_BRANCH% ./..."},
		{ID: 2, Name: "Name2", Path: "tool2"},
		{ID: 3, Name: "Name2", Path: "tool3"},
	mockDB := db.NewMockDB()
	analysis, _ := mockDB.StartAnalysis(1, 2)

	err := Analyse(context.Background(), analyser, tools, cfg, analysis)
	want := map[db.ToolID][]db.Issue{
		1: []db.Issue{{Path: "main.go", Line: 1, HunkPos: 1, Issue: "Name1: error1"}},
		2: []db.Issue{{Path: "main.go", Line: 1, HunkPos: 1, Issue: "Name2: error2"}},
		3: nil,
	}
	for toolID, issues := range want {
		if have := analysis.Tools[toolID].Issues; !reflect.DeepEqual(issues, have) {
			t.Errorf("unexpected issues for toolID %v\nwant: %+v\nhave: %+v", toolID, issues, have)
		}
	if len(analysis.Tools) != len(want) {
		t.Errorf("analysis has %v tools want %v", len(analysis.Tools), len(want))
		{ID: 1, Name: "Name1", Path: "tool1", Args: "-flag %BASE_BRANCH% ./..."},
		{ID: 2, Name: "Name2", Path: "tool2"},
		{ID: 3, Name: "Name2", Path: "tool3"},
	mockDB := db.NewMockDB()
	analysis, _ := mockDB.StartAnalysis(1, 2)

	err := Analyse(context.Background(), analyser, tools, cfg, analysis)
	want := map[db.ToolID][]db.Issue{
		1: []db.Issue{{Path: "main.go", Line: 1, HunkPos: 1, Issue: "Name1: error1"}},
		2: []db.Issue{{Path: "main.go", Line: 1, HunkPos: 1, Issue: "Name2: error2"}},
		3: nil,
	for toolID, issues := range want {
		if have := analysis.Tools[toolID].Issues; !reflect.DeepEqual(issues, have) {
			t.Errorf("unexpected issues for toolID %v\nwant: %+v\nhave: %+v", toolID, issues, have)
		}
	}
	if len(analysis.Tools) != len(want) {
		t.Errorf("analysis has %v tools want %v", len(analysis.Tools), len(want))

	mockDB := db.NewMockDB()
	analysis, _ := mockDB.StartAnalysis(1, 2)

	err := Analyse(context.Background(), analyser, nil, cfg, analysis)