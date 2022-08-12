package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type StateType struct {
	Save    Save
	Windows fyne.Window
	Tabs    *container.AppTabs
}

type Save struct {
	MetaData MetaData
	UserData UserData
}

type MetaData struct {
	SavePool   savePool `json:"savePool"`
	CustomName struct {
		CustomName             string        `json:"customName"`
		Parameters             []interface{} `json:"parameters"`
		IsNameFromLocalization bool          `json:"isNameFromLocalization"`
	} `json:"customName"`
	SaveTime                  string  `json:"saveTime"`
	Version                   string  `json:"version"`
	SavedStateTypeName        string  `json:"savedStateTypeName"`
	Day                       int     `json:"day"`
	SecondsPlayed             float64 `json:"secondsPlayed"`
	CurrentLvl                int     `json:"currentLvl"`
	PotionsBrewed             int     `json:"potionsBrewed"`
	LegendarySubstancesBrewed int     `json:"legendarySubstancesBrewed"`
	ClientsServed             int     `json:"clientsServed"`
	Popularity                int     `json:"popularity"`
	Karma                     int     `json:"karma"`
}

type Inventory struct {
	Name          string `json:"name"`
	Count         int    `json:"count"`
	ClassFullName string `json:"classFullName"`
	Data          string `json:"data"`
}

type UserData struct {
	SystemsData []struct {
		SystemName string `json:"systemName"`
		Data       string `json:"data"`
	} `json:"systemsData"`
	CurrentLevel                    int  `json:"currentLevel"`
	LegendarySubstancesBrewedAmount int  `json:"legendarySubstancesBrewedAmount"`
	IsAnimating                     bool `json:"isAnimating"`
	Slots                           struct {
		MKeys   []string `json:"m_keys"`
		MValues []struct {
			Name          string `json:"name"`
			Count         int    `json:"count"`
			ClassFullName string `json:"classFullName"`
			Data          string `json:"data"`
		} `json:"m_values"`
	} `json:"slots"`
	UnlockedLegendaryRecipes []string `json:"unlockedLegendaryRecipes"`
	LegendaryRecipeBookmarks []struct {
		SerializedRails []struct {
			SerializedBookmarks []struct {
				Position struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"position"`
				PrefabIndex int  `json:"prefabIndex"`
				IsMirrored  bool `json:"isMirrored"`
			} `json:"serializedBookmarks"`
		} `json:"serializedRails"`
	} `json:"legendaryRecipeBookmarks"`
	LegendaryRecipes                          []string `json:"legendaryRecipes"`
	BrewedLegendaryRecipes                    []string `json:"brewedLegendaryRecipes"`
	ReadUnlockedLegendaryRecipes              []string `json:"readUnlockedLegendaryRecipes"`
	CurrentLegendaryRecipeIndex               int      `json:"currentLegendaryRecipeIndex"`
	WithSpawnedProduct                        bool     `json:"withSpawnedProduct"`
	UnlockedLegendaryRecipesNotificationCount int      `json:"unlockedLegendaryRecipesNotificationCount"`
	Day                                       int      `json:"day"`
	CanTakeIngredients                        bool     `json:"canTakeIngredients"`
	Heat                                      float64  `json:"heat"`
	PreviousHeat                              float64  `json:"previousHeat"`
	CoalsHeat                                 struct {
		MKeys   []string  `json:"m_keys"`
		MValues []float64 `json:"m_values"`
	} `json:"coalsHeat"`
	SavedColors []struct {
		Color struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"color"`
		Hsv struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"hsv"`
		IsDefaultColor bool `json:"isDefaultColor"`
	} `json:"savedColors"`
	ColorsHistory []struct {
		Color struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"color"`
		Hsv struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"hsv"`
		IsDefaultColor bool `json:"isDefaultColor"`
	} `json:"colorsHistory"`
	CurrentSpaceIndexForColor int `json:"currentSpaceIndexForColor"`
	SerializedCurrentPotion   struct {
		PotionUsedComponents []interface{} `json:"potionUsedComponents"`
		RecipeMarks          []struct {
			Type        int     `json:"type"`
			FloatValue  float64 `json:"floatValue"`
			StringValue string  `json:"stringValue"`
			Note        struct {
				NoteIndex int `json:"noteIndex"`
				NoteSize  int `json:"noteSize"`
				NoteDigit int `json:"noteDigit"`
			} `json:"note"`
		} `json:"recipeMarks"`
		CollectedPotionEffects []interface{} `json:"collectedPotionEffects"`
		PotionSkinSettings     struct {
			CurrentBottleName   string `json:"currentBottleName"`
			CurrentStickerName  string `json:"currentStickerName"`
			CurrentStickerAngle int    `json:"currentStickerAngle"`
			CurrentIconName     string `json:"currentIconName"`
			CurrentIconColor1   struct {
				R float64 `json:"r"`
				G float64 `json:"g"`
				B float64 `json:"b"`
				A float64 `json:"a"`
			} `json:"currentIconColor1"`
			CurrentIconColor2 struct {
				R float64 `json:"r"`
				G float64 `json:"g"`
				B float64 `json:"b"`
				A float64 `json:"a"`
			} `json:"currentIconColor2"`
			CurrentIconColor3 struct {
				R float64 `json:"r"`
				G float64 `json:"g"`
				B float64 `json:"b"`
				A float64 `json:"a"`
			} `json:"currentIconColor3"`
			CurrentIconColor4 struct {
				R float64 `json:"r"`
				G float64 `json:"g"`
				B float64 `json:"b"`
				A float64 `json:"a"`
			} `json:"currentIconColor4"`
			CurrentCustomTitle        string `json:"currentCustomTitle"`
			IsCurrentTitleCustom      bool   `json:"isCurrentTitleCustom"`
			CurrentDescription        string `json:"currentDescription"`
			CurrentRecipeNotes        string `json:"currentRecipeNotes"`
			IsCurrentIconCustom       bool   `json:"isCurrentIconCustom"`
			IsCurrentIconColor1Custom bool   `json:"isCurrentIconColor1Custom"`
			IsCurrentIconColor2Custom bool   `json:"isCurrentIconColor2Custom"`
			IsCurrentIconColor3Custom bool   `json:"isCurrentIconColor3Custom"`
			IsCurrentIconColor4Custom bool   `json:"isCurrentIconColor4Custom"`
			ColorsCount               int    `json:"colorsCount"`
		} `json:"potionSkinSettings"`
		SerializedPath struct {
			FixedPathPoints        []interface{} `json:"fixedPathPoints"`
			IndicatorRotationValue float64       `json:"indicatorRotationValue"`
			GlobalCrossParameters  struct {
				SpriteIndex int     `json:"spriteIndex"`
				Rotation    float64 `json:"rotation"`
			} `json:"globalCrossParameters"`
			GlobalSpiralParameters struct {
				SpriteIndex int     `json:"spriteIndex"`
				Rotation    float64 `json:"rotation"`
			} `json:"globalSpiralParameters"`
			DeletedGraphicsSegments              float64 `json:"deletedGraphicsSegments"`
			SegmentLengthToDeleteFromEndPhysics  float64 `json:"segmentLengthToDeleteFromEndPhysics"`
			SegmentLengthToDeleteFromEndGraphics float64 `json:"segmentLengthToDeleteFromEndGraphics"`
			SegmentLengthToDeleteFromEndDots     float64 `json:"segmentLengthToDeleteFromEndDots"`
			SegmentLengthToDeletePhysics         float64 `json:"segmentLengthToDeletePhysics"`
			SegmentLengthToDeleteGraphics        float64 `json:"segmentLengthToDeleteGraphics"`
			IndicatorTargetPosition              struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"indicatorTargetPosition"`
			PathPosition struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"pathPosition"`
			IndicatorPosition struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"indicatorPosition"`
			FollowButtonTargetObjectPosition struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"followButtonTargetObjectPosition"`
			Health float64 `json:"health"`
		} `json:"serializedPath"`
	} `json:"serializedCurrentPotion"`
	SerializedDefaultBottleSettings struct {
		DefaultBottleWeakPotionName     string `json:"defaultBottleWeakPotionName"`
		DefaultStickerWeakPotionName    string `json:"defaultStickerWeakPotionName"`
		DefaultStickerAngleWeakPotion   int    `json:"defaultStickerAngleWeakPotion"`
		DefaultBottleNormalPotionName   string `json:"defaultBottleNormalPotionName"`
		DefaultStickerNormalPotionName  string `json:"defaultStickerNormalPotionName"`
		DefaultStickerAngleNormalPotion int    `json:"defaultStickerAngleNormalPotion"`
		DefaultBottleStrongPotionName   string `json:"defaultBottleStrongPotionName"`
		DefaultStickerStrongPotionName  string `json:"defaultStickerStrongPotionName"`
		DefaultStickerAngleStrongPotion int    `json:"defaultStickerAngleStrongPotion"`
	} `json:"serializedDefaultBottleSettings"`
	PotionChangedAfterSavingRecipe bool `json:"potionChangedAfterSavingRecipe"`
	PreviousPotionRecipe           struct {
		Name      string `json:"name"`
		ClassName string `json:"className"`
		Data      string `json:"data"`
	} `json:"previousPotionRecipe"`
	DialoguesGlobalConditionProperties []interface{} `json:"dialoguesGlobalConditionProperties"`
	DialoguesLocalConditionProperties  []interface{} `json:"dialoguesLocalConditionProperties"`
	CurrentDialoguePath                []interface{} `json:"currentDialoguePath"`
	ReadDialogueNodes                  []string      `json:"readDialogueNodes"`
	TraderTradingPanelInventory        []interface{} `json:"traderTradingPanelInventory"`
	PlayerTradingPanelInventory        []interface{} `json:"playerTradingPanelInventory"`
	ItemCostMultipliers                []interface{} `json:"itemCostMultipliers"`
	NextDialogueState                  int           `json:"nextDialogueState"`
	WindSpeed                          float64       `json:"windSpeed"`
	TotalDayNpcCount                   int           `json:"totalDayNpcCount"`
	ServedNpcCount                     int           `json:"servedNpcCount"`
	CurrentMortarStains                []interface{} `json:"currentMortarStains"`
	MortarStainsToRemove               []interface{} `json:"mortarStainsToRemove"`
	CurrentPestleStains                []interface{} `json:"currentPestleStains"`
	PestleStainsToRemove               []interface{} `json:"pestleStainsToRemove"`
	Experience                         struct {
		CurrentExp   float64 `json:"currentExp"`
		NextLvlAt    float64 `json:"nextLvlAt"`
		CurrentLvlAt float64 `json:"currentLvlAt"`
		CurrentLvl   int     `json:"currentLvl"`
	} `json:"experience"`
	ExperienceBonusMapItems []struct {
		LocalPosition struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"localPosition"`
		AlreadyCollected bool `json:"alreadyCollected"`
		IsDaily          bool `json:"isDaily"`
	} `json:"experienceBonusMapItems"`
	Chunks []struct {
		MapIndex       int `json:"mapIndex"`
		PositionInGrid struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"positionInGrid"`
		SumOfTilesFogDensity int   `json:"sumOfTilesFogDensity"`
		Tiles                []int `json:"tiles"`
	} `json:"chunks"`
	CurrentChapterIndex int `json:"currentChapterIndex"`
	GoalsBookBookmarks  []struct {
		SerializedRails []struct {
			SerializedBookmarks []struct {
				Position struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"position"`
				PrefabIndex int  `json:"prefabIndex"`
				IsMirrored  bool `json:"isMirrored"`
			} `json:"serializedBookmarks"`
		} `json:"serializedRails"`
	} `json:"goalsBookBookmarks"`
	ChaptersGroups []struct {
		Name     string `json:"name"`
		Chapters []struct {
			Name        string `json:"name"`
			ChapterGoal struct {
				Name             string `json:"name"`
				Progress         int    `json:"progress"`
				FollowState      int    `json:"followState"`
				IsCompletionRead bool   `json:"isCompletionRead"`
			} `json:"chapterGoal"`
			Goals []struct {
				Name             string `json:"name"`
				Progress         int    `json:"progress"`
				FollowState      int    `json:"followState"`
				IsCompletionRead bool   `json:"isCompletionRead"`
			} `json:"goals"`
		} `json:"chapters"`
	} `json:"chaptersGroups"`
	GrowingSpotsMain     []interface{} `json:"growingSpotsMain"`
	GrowingSpotsTutorial []struct {
		SpotName  string `json:"spotName"`
		PlantName string `json:"plantName"`
	} `json:"growingSpotsTutorial"`
	PointerState       int     `json:"pointerState"`
	PointerDirection   int     `json:"pointerDirection"`
	PointerPosition    float64 `json:"pointerPosition"`
	DifficultyToSelect int     `json:"difficultyToSelect"`
	CurrentDifficulty  int     `json:"currentDifficulty"`
	NpcBonuses         struct {
		MKeys   []int `json:"m_keys"`
		MValues []struct {
			Infos []struct {
				Size       float64 `json:"size"`
				ThemeTitle string  `json:"themeTitle"`
				BonusIndex int     `json:"bonusIndex"`
				SizeIndex  int     `json:"sizeIndex"`
				Mirrored   bool    `json:"mirrored"`
				Position   float64 `json:"position"`
			} `json:"infos"`
		} `json:"m_values"`
	} `json:"npcBonuses"`
	NpcHaggleThemes struct {
		MKeys   []int    `json:"m_keys"`
		MValues []string `json:"m_values"`
	} `json:"npcHaggleThemes"`
	NpcHaggleSeeds struct {
		MKeys   []int `json:"m_keys"`
		MValues []int `json:"m_values"`
	} `json:"npcHaggleSeeds"`
	CurrentBonuses []struct {
		Size       float64 `json:"size"`
		ThemeTitle string  `json:"themeTitle"`
		BonusIndex int     `json:"bonusIndex"`
		SizeIndex  int     `json:"sizeIndex"`
		Mirrored   bool    `json:"mirrored"`
		Position   float64 `json:"position"`
	} `json:"currentBonuses"`
	StateBeforeHaggleStarted int   `json:"stateBeforeHaggleStarted"`
	HaggleState              int   `json:"haggleState"`
	UnlockedDifficulties     []int `json:"unlockedDifficulties"`
	ItemsFromInventory       []struct {
		TypeName string `json:"typeName"`
		Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"position"`
		EulerAngles struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"eulerAngles"`
		InventoryItemName string `json:"inventoryItemName"`
		Data              string `json:"data"`
	} `json:"itemsFromInventory"`
	Karma             int           `json:"karma"`
	NpcQueue          []interface{} `json:"npcQueue"`
	SpawnedNpcQueue   []interface{} `json:"spawnedNpcQueue"`
	NpcGlobalSettings struct {
		MKeys   []string `json:"m_keys"`
		MValues []struct {
			Closeness                        int   `json:"closeness"`
			Cooldown                         int   `json:"cooldown"`
			NpcHasVisitedShopAtCloseness     []int `json:"npcHasVisitedShopAtCloseness"`
			CompletedUniqueQuestsAtCloseness []int `json:"completedUniqueQuestsAtCloseness"`
		} `json:"m_values"`
	} `json:"npcGlobalSettings"`
	CurrentNpc struct {
		Npc struct {
			NpcForSpawn struct {
				NpcTemplateName            string  `json:"npcTemplateName"`
				NpcFactionName             string  `json:"npcFactionName"`
				NpcFactionClassName        string  `json:"npcFactionClassName"`
				NpcMood                    int     `json:"npcMood"`
				ChapterOnAddToSpawn        int     `json:"chapterOnAddToSpawn"`
				PathTweenFullPosition      float64 `json:"pathTweenFullPosition"`
				PauseBeforeSpawn           bool    `json:"pauseBeforeSpawn"`
				PartsGenerationRandomState struct {
					S0 int `json:"s0"`
					S1 int `json:"s1"`
					S2 int `json:"s2"`
					S3 int `json:"s3"`
				} `json:"partsGenerationRandomState"`
				TradingRandomState struct {
					S0 int `json:"s0"`
					S1 int `json:"s1"`
					S2 int `json:"s2"`
					S3 int `json:"s3"`
				} `json:"tradingRandomState"`
				SpawnPosition struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"spawnPosition"`
			} `json:"npcForSpawn"`
			State                   int `json:"State"`
			QuestsOnCooldownOnSpawn struct {
				MKeys   []interface{} `json:"m_keys"`
				MValues []interface{} `json:"m_values"`
			} `json:"questsOnCooldownOnSpawn"`
		} `json:"npc"`
		IsPotionSold         bool `json:"isPotionSold"`
		IsHaggleCanceled     bool `json:"isHaggleCanceled"`
		IsCurrentDealHaggled bool `json:"isCurrentDealHaggled"`
		IsAnyDealHaggled     bool `json:"isAnyDealHaggled"`
		Gold                 int  `json:"gold"`
		Closeness            int  `json:"closeness"`
		Bargained            int  `json:"bargained"`
		RandomGreetingsIndex int  `json:"randomGreetingsIndex"`
	} `json:"currentNpc"`
	TutorialNpcQueueCount     int  `json:"tutorialNpcQueueCount"`
	WasFirstMainTraderSpawned bool `json:"wasFirstMainTraderSpawned"`
	KarmicTradersVirtualQueue struct {
		Queue struct {
			MKeys   []interface{} `json:"m_keys"`
			MValues []interface{} `json:"m_values"`
		} `json:"queue"`
		TemporaryPool []interface{} `json:"temporaryPool"`
	} `json:"karmicTradersVirtualQueue"`
	ExtraTradersVirtualQueue struct {
		Queue struct {
			MKeys   []int    `json:"m_keys"`
			MValues []string `json:"m_values"`
		} `json:"queue"`
		TemporaryPool []interface{} `json:"temporaryPool"`
	} `json:"extraTradersVirtualQueue"`
	MainTradersSpawnRepeatDefender struct {
		SpawnedNpc []string `json:"spawnedNpc"`
	} `json:"mainTradersSpawnRepeatDefender"`
	QuestsOnCooldown struct {
		MKeys   []string `json:"m_keys"`
		MValues []int    `json:"m_values"`
	} `json:"questsOnCooldown"`
	PestlePosition struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"pestlePosition"`
	PestleRotation struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"pestleRotation"`
	Gold                 int          `json:"gold"`
	GoldEarned           int          `json:"goldEarned"`
	ProfitFromHaggling   int          `json:"profitFromHaggling"`
	PlayerPanelType      int          `json:"playerPanelType"`
	PlayerInventory      []*Inventory `json:"playerInventory"`
	Popularity           int          `json:"popularity"`
	PotionEffectMapItems []struct {
		LocalPosition struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"localPosition"`
		Status int `json:"status"`
	} `json:"potionEffectMapItems"`
	SavedRecipes []struct {
		Name      string `json:"name"`
		ClassName string `json:"className"`
		Data      string `json:"data"`
	} `json:"savedRecipes"`
	CurrentRecipeIndex        int `json:"currentRecipeIndex"`
	CurrentPotionRecipeIndex  int `json:"currentPotionRecipeIndex"`
	UnlockedRecipesPagesCount int `json:"unlockedRecipesPagesCount"`
	RecipeBookBookmarks       []struct {
		SerializedRails []struct {
			SerializedBookmarks []struct {
				Position struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
				} `json:"position"`
				PrefabIndex int  `json:"prefabIndex"`
				IsMirrored  bool `json:"isMirrored"`
			} `json:"serializedBookmarks"`
		} `json:"serializedRails"`
	} `json:"recipeBookBookmarks"`
	Zoom        float64 `json:"zoom"`
	MapPosition struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"mapPosition"`
	FollowIndicatorButtonState bool     `json:"followIndicatorButtonState"`
	SelectedPotionBaseName     string   `json:"selectedPotionBaseName"`
	DefaultPotionBaseName      string   `json:"defaultPotionBaseName"`
	UnlockedPotionBases        []string `json:"unlockedPotionBases"`
	UnlockedPotionBasesRead    []bool   `json:"unlockedPotionBasesRead"`
	Vortexes                   []struct {
		IsLocked bool `json:"isLocked"`
		Position struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"position"`
		IsIndicatorMovingThroughVortex bool `json:"isIndicatorMovingThroughVortex"`
		PathShift                      struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pathShift"`
	} `json:"vortexes"`
	CurrentRoom                             int  `json:"currentRoom"`
	IsPlayerVisitedMeetingRoomAfterDayStart bool `json:"isPlayerVisitedMeetingRoomAfterDayStart"`
	Scales                                  struct {
		FluctuationTime                 float64 `json:"fluctuationTime"`
		Amplitude                       float64 `json:"amplitude"`
		Phase                           float64 `json:"phase"`
		AmplitudesSum                   float64 `json:"amplitudesSum"`
		IsFluctuating                   bool    `json:"isFluctuating"`
		IsShaking                       bool    `json:"isShaking"`
		IsShakingToTheRight             bool    `json:"isShakingToTheRight"`
		MovingSideSign                  int     `json:"movingSideSign"`
		IsCorrectBonusClicked           bool    `json:"isCorrectBonusClicked"`
		IsIncorrectBonusClicked         bool    `json:"isIncorrectBonusClicked"`
		TimeCoefficient                 float64 `json:"timeCoefficient"`
		IsLowestPositionOnHaggleReached bool    `json:"isLowestPositionOnHaggleReached"`
		IsWrongPotionOnTheScales        bool    `json:"isWrongPotionOnTheScales"`
		CapAngleOnReturnFromMaxAngle    bool    `json:"capAngleOnReturnFromMaxAngle"`
		CurrentAngle                    float64 `json:"currentAngle"`
		TargetAngle                     float64 `json:"targetAngle"`
		EnablePhysics                   bool    `json:"enablePhysics"`
		CurrentState                    int     `json:"currentState"`
	} `json:"scales"`
	ScalesPotionItem struct {
		Name          string `json:"name"`
		Count         int    `json:"count"`
		ClassFullName string `json:"classFullName"`
		Data          string `json:"data"`
	} `json:"scalesPotionItem"`
	IsCurrentPotionSuitable bool `json:"isCurrentPotionSuitable"`
	IsPotionItemOnScales    bool `json:"isPotionItemOnScales"`
	SpoonPosition           struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"spoonPosition"`
	SpoonRotation struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	} `json:"spoonRotation"`
	SpoonAnimationFrameNumber int `json:"spoonAnimationFrameNumber"`
	SpoonDirtyPartsList       []struct {
		MaskLocalPosition struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"maskLocalPosition"`
		PartColor struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
			A float64 `json:"a"`
		} `json:"partColor"`
		ContaminationValue float64 `json:"contaminationValue"`
	} `json:"spoonDirtyPartsList"`
	SecondsPlayed float64 `json:"secondsPlayed"`
	ClientsServed int     `json:"clientsServed"`
	PotionsBrewed int     `json:"potionsBrewed"`
	Talents       struct {
		CurrentPoints      int      `json:"currentPoints"`
		EarnedTalentsNames []string `json:"earnedTalentsNames"`
	} `json:"talents"`
	ShowTalentsNotification bool          `json:"showTalentsNotification"`
	TraderPanelType         int           `json:"traderPanelType"`
	TraderInventory         []interface{} `json:"traderInventory"`
	TutorialState           struct {
		SetName                string        `json:"setName"`
		StepIndex              int           `json:"stepIndex"`
		ManipulatedObjectsLock []interface{} `json:"manipulatedObjectsLock"`
		ResetState             string        `json:"resetState"`
	} `json:"tutorialState"`
}
