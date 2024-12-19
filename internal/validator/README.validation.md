### why put the validation in domain layer?

- Is it presentation logic or domain logic? Presentation logic is something you decide "mapping render model", "format
  of render model", "how to render", "what color, what size, which text", "how long will it stay on screen" etc... If
  validation is presentation logic, why does backend code have same validation control? From my perspective, validation
  is Domain logic.


- Why validation is Domain logic? Who decides if username can be 20 char at max? Business rule decides. Who decides
  number of max items in shopping basket? Business rule decides. The length of username is decision of business, and
  that rule
  applies in everywhere in the project. CreateProfile/ UpdateProfile/ Register etc.. all have same max-20char-username
  rule. That length control (validation) code should reside in Domain layer.


- What is the flow if validation code is in Domain layer? User clicks button in View. ViewModel/Presenter calls domain
  layer function. Domain layer function validates input data. If there are invalid input parameters, it returns
  ValidationException with explanation. ValidationException will contain list of invalid parameters, type of validation
  they failed (minLength, maxLength, emailPatternMismatch etc..), what is expected (20 char at max etc..).
  ViewModel/Presenter/Controller gets this ValidationException and here we have Presentation logic. Now it decides what
  to render, how to render. Do we render error of all invalid inputs or only first invalid input? What text/color should
  be shown (based on data in ValidationException) ?
  Do we render error as popup/textView/tooltip? After all presentation
  decisions are made and new model is created, View just! renders using that model.


- Another point is, in Domain layer, where should be validation code? In UseCase functions or in Models (why not)
  itself?
  IMHO, there should be Stateless Generic Interface/Class that has generic validation logics. And after that point, each
  UseCase class can implement ValidationInterface or inject it as Class object. If multiple UseCases need same
  validation, validation control logic will be duplicated. What happens if we put validation logic in Model itself?
  Model would implement ValidationInterface (which has stateless pure functions only!) and have fun validate():
  ValidationOutcome function.
  I don't think it is problem to put validation logic of Business Model in itself. All UseCases would call
  model.validate() only. There is dependency between Model and ValidationOutcome.