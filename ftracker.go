package ftracker

import (
	"fmt"
	"math"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                            = 0.65  // Средняя длина шага.
	mInKm                              = 1000  // Количество метров в километре.
	minInH                             = 60    // Количество минут в часе.
	kmhInMsec                          = 0.278 // Коэффициент для преобразования км/ч в м/с.
	cmInM                              = 100   // Количество сантиметров в метре.
	runningCaloriesMeanSpeedMultiplier = 18    // Множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 1.79  // Среднее количество сжигаемых калорий при беге.
	walkingCaloriesWeightMultiplier    = 0.035 // Множитель массы тела.
	walkingSpeedHeightMultiplier       = 0.029 // Множитель роста.
	swimmingCaloriesMeanSpeedShift     = 1.1   // Среднее количество сжигаемых калорий при плавании относительно скорости.
	swimmingCaloriesWeightMultiplier   = 2     // Множитель веса при плавании.
)

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	dist := distance(action)
	return dist / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch trainingType {
	case "Бег":
		dist := distance(action)
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	case "Ходьба":
		dist := distance(action)
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	case "Плавание":
		dist := float64(lengthPool*countPool) / mInKm
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// RunningSpentCalories возвращает количество потраченных калорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	speed := meanSpeed(action, duration)
	return ((runningCaloriesMeanSpeedMultiplier*speed + runningCaloriesMeanSpeedShift) * weight) * duration * minInH
}

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	speed := meanSpeed(action, duration) / kmhInMsec
	heightM := height / cmInM
	return (walkingCaloriesWeightMultiplier*weight + (math.Pow(speed, 2)/heightM)*walkingSpeedHeightMultiplier*weight) * duration * minInH
}

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool) * float64(countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
