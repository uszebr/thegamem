package modelfactory

import (
	"sync"

	"github.com/uszebr/thegamem/play/model"
	"github.com/uszebr/thegamem/play/model/implementation/fiveinteractions/modfivestepfinalconclusion"
	"github.com/uszebr/thegamem/play/model/implementation/greenexploit"
	"github.com/uszebr/thegamem/play/model/implementation/modalwaysgreen"
	"github.com/uszebr/thegamem/play/model/implementation/modalwaysred"
	"github.com/uszebr/thegamem/play/model/implementation/modblindrevenge"
	"github.com/uszebr/thegamem/play/model/implementation/modcopystartgreen"
	"github.com/uszebr/thegamem/play/model/implementation/modcopystartred"
	"github.com/uszebr/thegamem/play/model/implementation/modrandom"
)

var (
	modalFactory ModelFactory
	once         sync.Once
)

type ModelCreatorI interface {
	GetModel() model.Model
}

type ModelFactory struct {
	models     map[string]ModelCreatorI
	modelNames []string
}

// Singleton
func GetModelFactory() *ModelFactory {
	once.Do(func() {
		models := map[string]ModelCreatorI{
			"alwaysgreen":             modalwaysgreen.ModAlwaysGreen{},
			"alwaysred":               modalwaysred.ModAlwaysRed{},
			"blindrevenge":            modblindrevenge.ModBlindRevenge{},
			"copystartgreen":          modcopystartgreen.ModCopyStrartGreen{},
			"copystartred":            modcopystartred.ModCopyStrartRed{},
			"random":                  modrandom.ModRandom{},
			"fivestepfinalconclusion": modfivestepfinalconclusion.Modfivestepfinalconclusion{},
			"greenexploit":            greenexploit.ModGreenExploit{},
		}
		modelNames := make([]string, 0, len(models))
		for key := range models {
			modelNames = append(modelNames, key)
		}
		modalFactory = ModelFactory{models: models, modelNames: modelNames}
	})
	return &modalFactory
}

// not returns error, if model not found - panic!!!! Doesnt make sense to continue
func (factory ModelFactory) MustCreateModel(modelName string) model.Model {
	m, ok := factory.models[modelName]
	if !ok {
		panic("Can not get model for: " + modelName)
	}
	return m.GetModel()
}

func (factory ModelFactory) GetAllModelNames() []string {
	return factory.modelNames
}
