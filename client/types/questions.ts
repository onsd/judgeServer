export interface QuestionType {
    ID?:         number;
    CreatedAt?:  Date;
    UpdatedAt?:  Date;
    DeletedAt?:  null;
    title:      string;
    body:       string;
    validation: string;
    input:      string;
    output:     string;
    testcase:   Case[];
    samplecase: Case[];
}

export interface Case {
    ID?:         number;
    CreatedAt?:  Date;
    UpdatedAt?:  Date;
    DeletedAt?:  null;
    QuestionID?: number;
    Input:      string;
    Output:     string;
}

export interface AnswerType {
    ID?:          number;
    CreatedAt:   Date;
    UpdatedAt:   Date;
    DeletedAt:   null;
    user_id:     number;
    question_id: number;
    language:    string;
    answer:      string;
    status:      string;
    result:      string;
    detail:      string;
}
