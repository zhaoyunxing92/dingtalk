/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package priority

//Level 优先级
type Level int

//优先级，取值：
//10：较低
//20：普通
//30：紧急
//40：非常紧急
const (
	//Lower 较低
	Lower = Level(10)

	//Ordinary 普通
	Ordinary = Level(20)

	//Emergency 紧急
	Emergency = Level(30)

	//Urgent 非常紧急
	Urgent = Level(40)
)
